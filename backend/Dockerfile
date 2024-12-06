# Stage 1: Build the application
FROM golang:1.23.1-alpine as builder

# Install necessary dependencies in Alpine
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application binary
RUN go build -o tvapp-backend ./server.go

# Stage 2: Final image to run the application
FROM alpine:3.18

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/tvapp-backend .

# (Optional) Copy the .env file into the container
# Comment out this line if you plan to use Option 2 (mounting the .env file)
COPY .env .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./tvapp-backend"]
