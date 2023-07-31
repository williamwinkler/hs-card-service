# Start from a lightweight base image with Go support
FROM golang:1.19.4-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code into the container
COPY . .

# Build the Go binary inside the container
RUN CGO_ENABLED=0 GOOS=linux go build -o hs-card-service

# Use a minimal base image to create the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary built in the previous stage into the final image
COPY --from=builder /app/hs-card-service .

# Expose the port that the Go application listens on
EXPOSE 8080

# Command to run the Go application
CMD ["./hs-card-service"]
