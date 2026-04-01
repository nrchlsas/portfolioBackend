.PHONY: help install run build test clean db-up db-down swagger

help:
	@echo "Portfolio Backend - Available Commands"
	@echo ""
	@echo "  make install      - Install Go dependencies"
	@echo "  make run          - Run the development server"
	@echo "  make build        - Build the binary"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make db-up        - Start MySQL database using Docker Compose"
	@echo "  make db-down      - Stop MySQL database"
	@echo "  make swagger      - Generate Swagger documentation"
	@echo "  make dev          - Run in development mode with hot reload"
	@echo ""

install:
	go mod download
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest

run:
	go run main.go

build:
	go build -o bin/portfolio-backend main.go

test:
	go test -v ./...

clean:
	rm -rf bin/
	rm -rf docs/

db-up:
	docker-compose up -d

db-down:
	docker-compose down

swagger:
	swag init

dev:
	go run main.go

all: install db-up swagger run
