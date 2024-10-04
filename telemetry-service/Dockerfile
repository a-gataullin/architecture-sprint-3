# Use the official Golang image to build the service
FROM golang:1.20 as builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o telemetry-service ./cmd/main.go

# Use a minimal base image for the final build
FROM alpine:latest

# Copy the compiled binary from the builder stage
COPY --from=builder /app/telemetry-service /usr/local/bin/

# Set the entry point for the container
ENTRYPOINT ["/usr/local/bin/telemetry-service"]

# Expose the port the service will run on
EXPOSE 8081
