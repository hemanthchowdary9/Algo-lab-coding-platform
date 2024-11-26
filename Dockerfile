# Use the official Go image as a build environment
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to install dependencies
COPY go.mod ./
COPY go.sum ./

# Install dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o platform ./cmd/platform

# Use a lightweight image to run the app
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/platform .
COPY --from=builder /app/config/ ./config/

# Copy the static and template files
COPY ./templates ./templates
COPY ./static ./static

# Expose the application port
EXPOSE 7074

# Run the Go app
CMD ["./platform"]
