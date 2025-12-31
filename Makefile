# Makefile

AUTH_PROTO_DIR := ./proto/auth/v1
USERS_PROTO_DIR := ./proto/users/v1

AUTH_SERVICE_DIR := ./services/auth
USERS_SERVICE_DIR := ./services/users

# === Auth ===

.PHONY: proto-auth
proto-auth:
	protoc \
		--proto_path=./proto \
		--go_out=$(AUTH_SERVICE_DIR)/pkg/proto \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(AUTH_SERVICE_DIR)/pkg/proto \
		--go-grpc_opt=paths=source_relative \
		$(AUTH_PROTO_DIR)/auth.proto

.PHONY: mod-auth
mod-auth:
	cd $(AUTH_SERVICE_DIR) && go mod tidy

.PHONY: build-auth
build-auth:
	cd $(AUTH_SERVICE_DIR) && go build -o bin/auth ./cmd/app

.PHONY: run-auth
run-auth:
	cd $(AUTH_SERVICE_DIR) && go run ./cmd/app

# === Users ===

.PHONY: proto-users
proto-users:
	protoc \
		--proto_path=./proto \
		--go_out=$(USERS_SERVICE_DIR)/pkg/proto \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(USERS_SERVICE_DIR)/pkg/proto \
		--go-grpc_opt=paths=source_relative \
		$(USERS_PROTO_DIR)/users.proto

.PHONY: mod-users
mod-users:
	cd $(USERS_SERVICE_DIR) && go mod tidy

.PHONY: build-users
build-users:
	cd $(USERS_SERVICE_DIR) && go build -o bin/users ./cmd/app

.PHONY: run-users
run-users:
	cd $(USERS_SERVICE_DIR) && go run ./cmd/app

# === Общие цели ===

.PHONY: proto
proto: proto-auth proto-users

.PHONY: mod
mod: mod-auth mod-users

.PHONY: build
build: build-auth build-users

.PHONY: run
run: run-auth run-users

.PHONY: lint
lint:
	cd $(AUTH_SERVICE_DIR) && golangci-lint run ./...

# Сборка только auth (как у тебя было)
.PHONY: auth
auth: proto-auth mod-auth build-auth