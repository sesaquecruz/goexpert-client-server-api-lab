up-db:
	@./sqlite.sh

up-server: up-db
	@go run cmd/server/main.go

up-client:
	@go run cmd/client/main.go
