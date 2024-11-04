# Define variables
APP_NAME = blog-app
DOCKER_COMPOSE_FILE = docker-compose.yml
MIGRATION_FOLDER = migrations
DATABASE_URL = postgres://postgres:sairam123@localhost:5432/blog_app?sslmode=disable
DB_DOCKER_CONTAINER=blog_db_container
CONFIG_PATH = config/config.yaml


# Set the binary output folder
BINARY_PATH = ./bin

# make create-network
create-network:
	docker network create blog_network

# creating the container with postgres software
# make postgres
postgres:
	@echo "${DB_DOCKER_CONTAINER}"
	docker run --network blog_network --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=sairam123 -d postgres:12-alpine

# creating ecommerce_db database inside the postgres container
# make createdb 
createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=postgres --owner=postgres blog_app 

# make start-docker
start-docker:
	docker start ${DB_DOCKER_CONTAINER}



# Default command: build and run the application
.PHONY: all
all: build run

# Build the Go application
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	go build -o $(BINARY_PATH)/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

.PHONY: run_app
run_app:
	@echo "Running $(APP_NAME)..."
	./bin/$(APP_NAME) -config $(CONFIG_PATH)

# Run the application (Docker Compose)
.PHONY: run
run:
	@echo "Starting $(APP_NAME) using Docker Compose..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up --build

# Run the application in detached mode
.PHONY: rund
rund:
	@echo "Starting $(APP_NAME) in detached mode..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d --build

# Stop the application
.PHONY: stop
stop:
	@echo "Stopping $(APP_NAME)..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Clean up binaries and containers
.PHONY: clean
clean:
	@echo "Cleaning up..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --volumes --remove-orphans
	rm -f $(BINARY_PATH)/$(APP_NAME)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test ./... -v

# Run migrations (using Goose)
.PHONY: migrate
migrate:
	@echo "Running migrations..."
	goose -dir $(MIGRATION_FOLDER) postgres "$(DATABASE_URL)" up

# Rollback migrations (using Goose)
.PHONY: rollback
rollback:
	@echo "Rolling back last migration..."
	goose -dir $(MIGRATION_FOLDER) postgres "$(DATABASE_URL)" down

# Create a new migration file (requires name)
.PHONY: new-migration
new-migration:
ifndef name
	$(error "name is required. Usage: make new-migration name=your_migration_name")
endif
	@echo "Creating new migration: $(name)"
	goose -dir $(MIGRATION_FOLDER) create $(name) sql

# View migration status
.PHONY: migration-status
migration-status:
	@echo "Checking migration status..."
	goose -dir $(MIGRATION_FOLDER) postgres "$(DATABASE_URL)" status
