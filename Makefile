.SILENT:
.PHONY: up down lint

up:
	docker network create --driver bridge coinclientserver &> /dev/null || true
	docker-compose up -d

down:
	docker-compose down --rmi local
	docker network rm coinclientserver -f &> /dev/null || true

lint:
	golangci-lint run ./...
