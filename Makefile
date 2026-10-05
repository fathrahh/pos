.PHONY: dev dev-api dev-web build build-api build-web test tidy

dev-api:
	cd services/api && go run ./cmd/api

dev-web:
	pnpm --filter @pos-saas/web dev

dev:
	pnpm turbo run dev

build-api:
	cd services/api && go build -o bin/api ./cmd/api

build-web:
	pnpm --filter @pos-saas/web build

build: build-api build-web

test:
	cd services/api && go test -v -race ./...

tidy:
	cd services/api && go mod tidy
