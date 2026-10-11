# AGENTS.md — SaaS POS Monorepo Guidelines

This document defines architecture, codebase conventions, security guardrails, and coding standards for all AI agents and developers working on this repository.

---

## 1. Project Overview & Scope

- **Domain:** SaaS Point of Sale (POS) with Multi-Tenancy.
- **Phase 1 Goal:** Simple POS: Product catalog with COGS (HPP), multi-outlet inventory, fast transaction recording (cash/manual payments, architecture ready for payment gateways/QRIS), and gross/net profit metrics dashboard.
- **Repository Structure:** Polyglot Monorepo.

```
pos-saas/
├── apps/
│   ├── web/               # Svelte 5 (SPA mode via SvelteKit adapter-static)
│   └── mobile/            # React Native (Pure POS cashier app - Phase 2)
├── services/
│   └── api/               # Golang Clean Architecture REST API
├── go.work                # Go workspace orchestration
├── package.json           # Root package.json (pnpm workspaces)
├── pnpm-workspace.yaml    # Monorepo workspaces definition
├── turbo.json             # Turborepo task runner & cache config
├── Makefile               # Root orchestration for dev, test, build
└── .gitignore
```

---

## 2. Technology Stack

| Layer | Technology | Key Constraints |
|---|---|---|
| **Backend** | Golang (1.23+) | Clean Architecture, Modular Monolith |
| **API Router** | `gofiber/fiber/v2` | High-throughput, Express-like routing |
| **Database & ORM** | PostgreSQL 16 + `gorm.io/gorm` | Tenant-scoped queries via GORM hooks |
| **Auth** | Stateless JWT | Claims: `user_id`, `tenant_id`, `role`, `outlet_id` |
| **Frontend** | Svelte 5 (Runes) + SvelteKit | `adapter-static` (SPA: `ssr = false`, `fallback = 'index.html'`) |
| **Monorepo Tools** | `pnpm` + `Turborepo` + `go.work` + `Makefile` | No Nx allowed |

---

## 3. Multi-Tenancy & Authorization Model

### 3.1 Hierarchy
```
Tenant (Company / Business Entity)
  └── Outlets (Branches / Physical Stores)
        ├── Inventory (Stock tracked per outlet)
        └── Cashiers / Registers (Device / Shift level)
```

### 3.2 Database Schema Rules
Every tenant-scoped table MUST include:
- `tenant_id UUID NOT NULL` (Foreign key to `tenants.id`, indexed)
- Tables scoped to specific branches MUST additionally include `outlet_id UUID NOT NULL`.

### 3.3 RBAC Roles
- **`owner`**: Full tenant access, subscription management, outlet creation, staff account creation, cross-outlet financial reports.
- **`manager`**: Outlet-scoped product/stock management, staff scheduling, outlet-level profit and loss reporting.
- **`cashier`**: POS terminal only (record sales, open/close shifts, print receipts, no margin/HPP visibility).

### 3.4 Application Boundaries: Web Dashboard vs POS Cashier Terminal
- **Web App (`apps/web`):**
  - **Scope:** Pure Backoffice & Management Dashboard.
  - **Target Persona:** Owner & Branch/Store Manager.
  - **Core Capabilities:** Manage master products & variants, HPP/COGS, tiered pricing, packages/bundles, discount vouchers, multi-outlet stock inventory, staff roles, and analytics.
  - **Strict Constraint:** **NO POS cashier terminal/shift keypad mode on Web.**
- **POS Cashier Terminal (`apps/mobile` or dedicated tablet app — Phase 2 Spec):**
  - **Device Pairing Model (Planned for Phase 2):**
    1. One-time Master Setup: Owner/Manager logs in with email/pass to select outlet & provision a persistent `device_token`.
    2. Staff Cashier Fast Switch: 4-6 digit numeric PIN lock screen for cashiers to start shift, record orders, and print receipts without exposing master business credentials.

### 3.5 GORM Multi-Tenant Leak Prevention (CRITICAL)
- **Rule:** Never execute raw `db.Find()` or `db.Where()` without tenant isolation.
- **Mechanism:** Implement GORM plugin / scope extracting `tenant_id` from `context.Context`. If `tenant_id` is missing in context during an isolated query, the transaction MUST abort with an error.

```go
// Example repository query pattern
func (r *productRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
    tenantID := middleware.GetTenantID(ctx)
    var prod domain.Product
    err := r.db.WithContext(ctx).
        Where("tenant_id = ? AND id = ?", tenantID, id).
        First(&prod).Error
    return &prod, err
}
```

---

## 4. Backend Architecture (Golang Clean Architecture)

Location: `services/api/`

```
services/api/
├── cmd/
│   └── api/               # Entrypoint (main.go, wire up dependencies)
├── internal/
│   ├── config/            # Env and app config
│   ├── domain/            # Entities, domain errors, value objects (Pure Go, NO GORM tags)
│   ├── usecase/           # Business logic interfaces and implementations
│   ├── repository/        # GORM persistence models and SQL queries
│   ├── handler/http/      # Chi HTTP routes, request validation, response serialization
│   └── middleware/        # Tenant extractor, JWT auth, CORS, logging, recovery
└── pkg/
    ├── response/          # Standard JSON response envelope: { success, data, error }
    └── validator/         # Request DTO validation
```

### Business Logic Guardrails:
1. **Profit & COGS (HPP) Calculation:**
   - Products have `cogs` (cost of goods sold / base cost) and `price` (retail price).
   - Order items MUST freeze the snapshot of `cogs` and `unit_price` at the exact moment of sale. Never calculate historical profit using current master product costs.
2. **Idempotent Transactions:**
   - Order creation endpoints MUST accept a client-generated `idempotency_key` (UUID) to prevent double billing on network retry.
3. **Payment Extensibility:**
   - Payment logic must implement a generic interface `PaymentProcessor`. Phase 1 implements `CashProcessor` and `ManualTransferProcessor`. Future QRIS/Payment Gateway integrations will implement the same interface.

---

## 5. Frontend Guidelines (Svelte 5 Dashboard)

Location: `apps/web/`

- **Svelte 5 Syntax:** Use Runes (`$state`, `$derived`, `$props`, `$effect`). Do not use deprecated Svelte 4 reactivity (`let count = 0; $: double = count * 2`).
- **Mode:** Pure Client-Side Application (SPA). Data fetching happens in client `onMount` or event handlers, communicating via standard REST JSON.
- **No FE Unit Tests:** Per project requirement, testing effort is focused on backend logic (money, transactions, tenant isolation).

---

## 6. Monorepo Orchestration Commands

All cross-service commands are standardized in root `Makefile`:

- `make dev`: Concurrently start Go backend (`air` for hot reload) and Svelte web dev server.
- `make build`: Build Go binary and bundle web SPA.
- `make test`: Run backend unit tests (`go test -race ./...`).
- `make lint`: Run `golangci-lint` and web lint.

---

## 7. AI Agent Guardrails & Invariants

1. **Repository Context Tracking:**
   - `AGENTS.md` is tracked in Git as the source of truth for architectural invariants across different AI coding tools.
2. **Financial Precision:**
   - In Go, monetary amounts (prices, discounts, taxes, COGS) MUST use integer cents / minor units (e.g. `int64` representing Rupiah/cents) or arbitrary precision decimals (`shopspring/decimal`). NEVER use `float32` or `float64` for money arithmetic.
3. **Mandatory Tests on Financial Rules:**
   - Any modifications to pricing calculation, discounts, profit margins, or stock movement MUST have accompanying Go unit tests in `usecase/` passing before finishing the task.
4. **Idempotency & Concurrency:**
   - Deduct stock inside database transactions with row-level locks (`SELECT ... FOR UPDATE` or atomic decrement `WHERE stock >= quantity`).
