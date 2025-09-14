# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp .

# Stage 2: Create a smaller image to run the app
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

COPY --from=builder /app/myapp .
# Expose the port the app will run on
EXPOSE 8080

# Command to run the Go binary
CMD ["./myapp"]
