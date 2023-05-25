DOCKER_TAG := latest

run:
	go run ./cmd/.
up:
	docker compose up
down:
	docker compose down
build:
	docker build --platform linux/amd64 -t taiti09/mah-api:${DOCKER_TAG} --target deploy ./
build-local: ## Build docker image to local development
	docker compose build --no-cache
logs: ## Tail docker compose logs
	docker compose logs -f