# Defines the base image used in image creation.
FROM golang:1.20.4-alpine AS builder

# Add label to the image
LABEL author="Mohammad Alwi Irfani"
LABEL email=alwi.irfani1927@gmail.com
LABEL github=github.com/alwi09

# Executes the apk command to update the Alpine Linux packages and install Git.
RUN apk update && apk add --no-cache git

# Sets the working directory.
WORKDIR /home

# Copies the go.mod and go.sum files from the local directory to the current working directory inside the image.
COPY go.mod go.sum ./

RUN go mod tidy

RUN go mod download

COPY . .

RUN go build -o todoApp ./main.go

# Defines the second base image for the final running container.
FROM alpine:3.15

WORKDIR /home

COPY --from=builder /home/todoApp .

# Copies the migrations directory from the builder image to the migrations directory inside the final container.
COPY --from=builder /home/db/migration ./db/migration

# Exposes the port on which the application will run.
EXPOSE 8080

CMD ["./todoApp"]