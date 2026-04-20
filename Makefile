.PHONY: test build env db-init db-drop test-coverage

# application
APP_NAME := task

# build
BUILD_DIR := ./bin
BUILD_PATH := ./cmd/task-tracker-cli/

# sqlite3 storage
SCHEMA_PATH := ./scripts/schema.sql
STORAGE_PATH := ./storage/storage.db

TEST_PACKAGES := $$(go list ./... | grep -v -E 'mocks')

test:
	@go test -cover ${TEST_PACKAGES}

test-cover:
	@go test -coverprofile=cov.out ${TEST_PACKAGES}

test-coverage: test-cover
	@go tool cover -html=cov.out

build: env db-init
	@export CGO_ENABLED=1
	@go build -o ${BUILD_DIR}/${APP_NAME} ${BUILD_PATH}
	@echo "Application built!"

env:
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
	fi

db-init:
	@if [ ! -f ${STORAGE_PATH} ]; then \
		mkdir -p $(dir ${STORAGE_PATH}); \
		if [ -f ${SCHEMA_PATH} ]; then \
			sqlite3 ${STORAGE_PATH} < ${SCHEMA_PATH}; \
			echo "sqlite3 storage initialized!"; \
		else \
			echo "${SCHEMA_PATH} not found!"; \
			exit 1; \
		fi \
	fi

db-drop:
	@rm -rf $(dir $(STORAGE_PATH))
	@echo "sqlite3 storage dropped!"
