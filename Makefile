# ===== CONFIG =====
BACKEND_DIR := backend
FRONTEND_DIR := frontend

# ===== TARGETS =====

.PHONY: all backend frontend dev stop clean

all: dev

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

# optional: stop (only works if you background processes manually)
stop:
	@echo "stop manually (ctrl+c)"

# cleanup (optional)
clean:
	cd $(BACKEND_DIR) && go clean