# Load environment variables from .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: help run build db-up db-down db-shell

help:
	@echo "Available commands:"
	@echo "  run      : Run the web application with live-reloading"
	@echo "  build    : Compile the web application"
	@echo "  db-up    : Start the database container (detached)"
	@echo "  db-down  : Stop and remove the database container"
	@echo "  db-shell : Open a MySQL shell inside the container"

run:
	@echo "Starting web application with air..."
	@air

build:
	@echo "Generating templ files and compiling Go application..."
	@templ generate
	@go build -o ./tmp/web-app ./cmd/web/main.go

db-up:
	@echo "Starting database container..."
	@docker-compose up -d

db-down:
	@echo "Stopping database container..."
	@docker-compose down

db-shell:
	@echo "Connecting to MySQL shell..."
	@docker-compose exec db mysql -uuser -ppassword widgets
