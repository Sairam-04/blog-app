# Use Golang image to run migrations
FROM golang:1.23.2-alpine AS migration

# Set the working directory
WORKDIR /app

# Install git and other dependencies
RUN apk add --no-cache git curl

# Install Goose for migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Copy migration files
COPY migrations /app/migrations

# Copy the .env file if needed (optional)
# COPY ../.env /app/.env 

# Run migrations on startup
CMD ["goose", "-dir", "/app/migrations", "postgres", "$DB_URL", "up"]
