# Build stage
FROM golang:1.23.2-alpine AS build

WORKDIR /app/backend

# Install required modules
COPY go.mod go.sum Makefile ./
RUN go mod download

# Copy the rest of the code including the config directory
COPY . .

# Run the build using Makefile
RUN CGO_ENABLED=0 GOOS=linux go build -o blog-app cmd/blog-app/main.go

# Run time Stage
FROM alpine:latest

# Set working directory in the final image
WORKDIR /app/backend

# Copy the binary and the config file from the build stage
COPY --from=build /app/backend/blog-app .
COPY --from=build /app/backend/config/config.yaml ./config/

# Expose the application port
EXPOSE 8002

# Command to run the binary
CMD ["./blog-app", "-config=config/config.yaml"]
