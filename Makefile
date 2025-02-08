.SILENT:
.PHONY: up down lint

up:
	docker-compose up -d

down:
	docker-compose down --rmi local

lint:
	golangci-lint run ./...