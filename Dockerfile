# Use an official Golang runtime as a parent image
FROM golang:1.21-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Start a new stage from scratch
FROM scratch

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Define build arguments
ARG DB_DSN
ARG DB_HOST
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME
ARG DB_PORT

# Generate the .env file using build arguments
RUN echo "DB_DSN=${DB_DSN}" >> .env && \
    echo "DB_HOST=${DB_HOST}" >> .env && \
    echo "DB_USER=${DB_USER}" >> .env && \
    echo "DB_PASSWORD=${DB_PASSWORD}" >> .env && \
    echo "DB_NAME=${DB_NAME}" >> .env && \
    echo "DB_PORT=${DB_PORT}" >> .env

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
