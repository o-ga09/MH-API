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


# curl -X POST -H "Content-Type: application/json" -d '{"req": [{"name": "value1","desc": "value2","location": "value3","specify": "value4","weakness_attack": "value5","weakness_element": "value6"},{"name": "value7","desc": "value8","location": "value9","specify": "value10","weakness_attack": "value11","weakness_element": "value12"}]}' http://localhost:8080/v1/monsters/json
# curl -X DELETE -b "token='token'" http://localhost:8080/v1/auth/monsters/7
# curl -X POST -H 'Content-type: application/json' -d '{"name":"admin","password":"password"}' http://localhost:8080/v1/auth