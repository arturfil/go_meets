include .env

up: 
	docker-compose up -d --remove-orphans

down:
	docker-compose down

build.dev:
	go build -o ${BINARY} ./cmd/server/main.go

build.prod:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BINARY} ./cmd/server/main.go

run: build.dev
	@echo "Starting backend"
	@env PORT=${PORT} DSN=${DSN} JWT_SECRET=${JWT_SECRET} ./${BINARY} &
	@echo "Backend started"

start: up run

stop:
	@-pkill -SIGTERM -f "./${BINARY}"
	@echo "server stopped..."

test.unittests:
	go test -v --tags=unittests ./...

db.status:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(DSN) goose -dir=$(MIGRATIONS_PATH) status

db.up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DSN) goose -dir=$(MIGRATIONS_PATH) up

db.down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DSN) goose -dir=$(MIGRATIONS_PATH) down 

restart: stop run
