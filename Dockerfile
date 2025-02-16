# # Шаг 1: Базовый образ для сборки
# FROM golang:1.22 AS builder

# WORKDIR /app

# COPY go.mod go.sum ./

# RUN go mod download

# COPY . .

# RUN go build -o myapp .

# FROM debian:bullseye-slim

# WORKDIR /app

# COPY --from=builder /app/myapp .

# EXPOSE 8080

# CMD ["./myapp"]

FROM golang:alpine
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o avito-shop . 
EXPOSE 8080
CMD ["./avito-shop"]
# FROM golang:latest as build

# WORKDIR /app

# # Copy the Go module files
# COPY go.mod .
# COPY go.sum .

# # Download the Go module dependencies
# RUN go mod download

# COPY . .

# RUN go build -o /myapp .
 
# FROM alpine:latest as run

# # Copy the application executable from the build image
# COPY --from=build /myapp /myapp

# WORKDIR /app
# EXPOSE 8080
# CMD ["/myapp"]