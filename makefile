.PHONY: lint
lint:
	go vet ./...
	golangci-lint run

.PHONY: test
test:
	docker compose up -d mh-api-dbsrv01
	go test -parallel 1 ./...

.PHONY: test-coverage
test-coverage:
	go test ./... -coverprofile=coverage.out

.PHONY: generate
generate:
	go generate ./...

# DATABASE_URL is expected to be set in the environment
ifndef DATABASE_URL
$(warning DATABASE_URL is not set. Using default local development settings)
DATABASE_URL ?= mh-api:P@ssw0rd@tcp(127.0.0.1:3306)/mh-api?charset=utf8&parseTime=True&loc=Local
export DATABASE_URL
endif

.PHONY: migrate-up
migrate-up:
	$(eval MIGRATE_CMD := go run cmd/migration/main.go)
	$(MIGRATE_CMD) -command up

.PHONY: migrate-down
migrate-down:
	$(eval MIGRATE_CMD := go run cmd/migration/main.go)
	$(MIGRATE_CMD) -command down

.PHONY: migrate-status
migrate-status:
	$(eval MIGRATE_CMD := go run cmd/migration/main.go)
	$(MIGRATE_CMD) -command status

.PHONY: migrate-new
migrate-new:
	$(eval MIGRATE_CMD := go run cmd/migration/main.go)
	$(MIGRATE_CMD) -command new -name $(name)

.PHONY: seed
seed:
	$(eval MIGRATE_CMD := go run cmd/migration/main.go)
	$(MIGRATE_CMD) -command seed

.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: build
build:
	go build -o bin/api cmd/api/main.go

.PHONY: build-agent
build-agent:
	go build -o bin/agent cmd/agent/main.go

.PHONY: run-agent
run-agent:
	go run cmd/agent/main.go

.PHONY: docker-build-agent
docker-build-agent:
	docker build -t monhun-agent --target deploy-agent .

.PHONY: compose-up
compose-up:
	docker compose up -d mh-api-dbsrv01

.PHONY: compose-down
compose-down:
	docker compose down -v

.PHONY: compose-stop
compose-stop:
	docker compose down

.PHONY: compose-build
compose-build:
	docker compose up -d --build

.PHONY: compose-start
compose-start:
	docker compose up -d

.PHONY: e2e-prod
e2e-prod:
	scenarigo run --config e2e/config/scenarigo-prod.yaml

# Encode Image to Base64
.PHONY: encode-image
encode-image:
	$(eval BASE64_IMAGE := $(shell base64 -w 0 $(IMAGE_PATH)))
	$(eval export BASE64_IMAGE)

.PHONY: check
check: lint test test-coverage migrate-up migrate-status migrate-down seed run build compose-up compose-down compose-build compose-start encode-image compose-stop
