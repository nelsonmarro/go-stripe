STRIPE_SECRET=
STRIPE_KEY=

.PHONY: help run build

help:
	@echo "Available commands:"
	@echo "  run    : Run the web application with live-reloading"
	@echo "  build  : Compile the web application"

run:
	@echo "Starting web application with air..."
	@STRIPE_KEY=${STRIPE_KEY} STRIPE_SECRET=${STRIPE_SECRET} air

build:
	@echo "Generating templ files and compiling Go application..."
	@templ generate
	@go build -o ./tmp/web-app ./cmd/web/main.go
