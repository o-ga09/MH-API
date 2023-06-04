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
rm:
	docker compose down all --volumes --remove-orphans
test:
	go test -race ./...
post:
	curl -X POST http://localhost:8080/v1/monsters --data-urlencode 'name=リオレウス' --data-urlencode 'desc=空の王者' --data-urlencode 'location=平原' --data-urlencode 'specify=飛竜種' --data-urlencode 'weakness_A=頭部' --data-urlencode 'weakness_E=龍'