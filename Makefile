# ===== CONFIG =====
BACKEND_DIR := backend
FRONTEND_DIR := frontend

# ===== TARGETS =====
.PHONY: all backend frontend dev install stop clean

all: dev

# install dependencies
install:
	@echo "installing frontend dependencies"
	cd $(FRONTEND_DIR) && pnpm install

# run both
dev:
	@echo "starting backend + frontend"
	@$(MAKE) -j2 backend frontend

# backend only
backend:
	cd $(BACKEND_DIR) && go run ./cmd/server

# frontend only
frontend:
	cd $(FRONTEND_DIR) && pnpm dev

# cleanup (optional)
clean:
	cd $(BACKEND_DIR) && go clean