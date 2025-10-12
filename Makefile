include .env

up:
	@echo "Starting containers..."
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stopping containers..."
	docker-compose down

build:
	go build -o ${BINARY} ./api/

start:
	./${BINARY}

restart: build start