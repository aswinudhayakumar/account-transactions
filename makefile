dep:
	go mod tidy
	go mod vendor

build:
	docker build -t account-transactions:local .

run: build
	docker compose --file docker-compose.db.yml --file docker-compose.yml up

stop:
	docker stop account-transactions || true

down: 
	docker compose --file docker-compose.db.yml --file docker-compose.yml down

mock:
	go generate ./...

unit-test:
	go test -cover ./... -v