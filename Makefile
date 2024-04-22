include .env

up: 
	docker-compose up -d --remove-orphans

down:
	docker-compose down

build:
	go build -o ${BINARY} ./cmd/server/main.go

run: build
	@echo "Starting backend"
	@env PORT=${PORT} DSN=${DSN} JWT_SECRET=${JWT_SECRET} ./${BINARY} &
	@echo "Backend started"

start: up run

stop:
	@-pkill -SIGTERM -f "./${BINARY}"
	@echo "server stopped..."

test.unittests:
	go test -v --tags=unittests ./...

init.up:
	cat migrations/init.up.sql | docker exec -i ${DOCKER_CONTAINER_DB_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

init.down:
	cat migrations/init.down.sql | docker exec -i ${DOCKER_CONTAINER_DB_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

seed.up:
	cat migrations/seed.up.sql | docker exec -i ${DOCKER_CONTAINER_DB_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

seed.down:
	cat migrations/seed.down.sql | docker exec -i ${DOCKER_CONTAINER_DB_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

restart: stop run
