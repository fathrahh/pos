# POS SaaS (Monorepo)

A production-grade, multi-tenant Point of Sale (POS) application built with Golang and Svelte.

## 1. Tech Stack

- **Backend (`services/api`)**: Golang Clean Architecture, `chi` router, GORM + PostgreSQL.
- **Frontend Dashboard (`apps/web`)**: Svelte 5 (Runes) + Vite SPA mode.
- **Mobile POS (`apps/mobile`)**: React Native (planned Phase 2).
- **Monorepo Tooling**: `pnpm` workspaces, Turborepo, `go.work`, and root `Makefile`.

---

## 2. Directory Structure

```
pos-saas/
├── apps/
│   └── web/               # Svelte 5 Admin & Analytics Dashboard
├── services/
│   └── api/               # Golang REST API backend
├── go.work                # Go workspace configuration
├── package.json           # Root workspace scripts
├── pnpm-workspace.yaml    # PNPM packages mapping
├── turbo.json             # Turborepo task pipeline
└── Makefile               # Unified development commands
```

---

## 3. Quick Start

### Prerequisites
- [Go](https://go.dev/) (>= 1.23)
- [Node.js](https://nodejs.org/) (>= 20)
- [pnpm](https://pnpm.io/) (>= 8)

### Installation
```bash
# Install frontend & monorepo tools
pnpm install

# Tidy backend dependencies
make tidy
```

### Running Locally
```bash
# Run backend API (default port: 8080)
make dev-api

# Run web dashboard (default port: 3000)
make dev-web
```

### Building & Testing
```bash
# Build both backend binary and frontend static assets
make build

# Run backend tests
make test
```

---

## 4. API Endpoints (Starter)

| Method | Endpoint | Description | Headers |
|---|---|---|---|
| `GET` | `/health` | Health check probe | None |
| `GET` | `/api/v1/` | API root check | None |
| `GET` | `/api/v1/me` | Tenant context test | `X-Tenant-ID: <uuid>` |
