# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

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
# Define the build argument
ARG BASE_URL
ARG API_KEY
ARG API_SECRET

# Write the environment variables to a single file
RUN echo "BASE_URL=${BASE_URL}" > /root/app.env && \
    echo "API_KEY=${API_KEY}" >> /root/app.env && \
    echo "API_SECRET=${API_SECRET}" >> /root/app.env
# Copy the Go binary from the builder stage
COPY --from=builder /app/myapp .
# Expose the port the app will run on
EXPOSE 8080

# Command to run the Go binary
CMD ["./myapp"]
