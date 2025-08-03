# مرحله 1: Build app
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o auth-service

# مرحله 2: اجرای باینری داخل container سبک‌تر
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/auth-service .

EXPOSE 8080

CMD ["./auth-service"]
