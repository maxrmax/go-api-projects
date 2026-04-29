# go-api-project

Demo project with a Go backend and Nuxt frontend.

## Requirements

- Go 1.21+
- Node.js 18+
- pnpm

## Setup

```bash
make install
```

## Run both

```bash
make dev
```

## Run Backend

```bash
cd backend
go run cmd/server/main.go
```

Runs on `http://localhost:8000`

## Run Frontend

```bash
cd frontend
pnpm install
pnpm dev
```

Runs on `http://localhost:3000`

## Default credentials

- Username: `admin`
- Password: `admin`