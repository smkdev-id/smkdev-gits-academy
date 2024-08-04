FROM golang:1.20.4-alpine AS builder

# Install build dependencies
RUN apk update && apk add --no-cache git gcc musl-dev

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app with main name is "main"
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Build the Docker image with the binary from the previous step
FROM alpine:latest  

# Add the certificates for the app
RUN apk --no-cache add ca-certificates

# Set the current working directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous step
COPY --from=builder /app/main .
COPY --from=builder /app/bookstore.db ./bookstore.db

# Expose port 5000 to the outside world
EXPOSE 5000

# Run the binary
CMD ["./main"]