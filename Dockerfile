# Use an official Go image as the base for building the application
FROM golang:1.23.4 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files and vendored dependencies
COPY go.mod go.sum vendor/ ./

# Copy the entire application code into the container
COPY . .

# Build the application binary
RUN CGO_ENABLED=0 GOOS=linux go build -o account-transactions ./cmd/account-transactions

# Use a minimal base image for running the application
FROM scratch

# Set the working directory inside the runtime container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/account-transactions .
COPY --from=builder /app/schema/migrations /schema/migrations 

# Expose the application port
EXPOSE 8080 8081

# Command to run the application
CMD ["./account-transactions"]
