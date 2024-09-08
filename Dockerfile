# Stage 1: Build the Go binary
FROM golang:1.20-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy all files to the working directory
COPY . .

# Build the Go application
RUN go mod tidy && go build -o key-server .

# Stage 2: Create a lightweight image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/key-server /app/key-server

# Expose the port the server will run on
EXPOSE 1123

# Command to run the application
CMD ["/app/key-server"]