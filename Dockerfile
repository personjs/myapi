FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/myapi cmd/main.go

FROM debian:bullseye

# Install necessary dependencies and upgrade glibc to version 2.34
RUN apt-get update && apt-get install -y ca-certificates curl gnupg2 lsb-release \
&& curl -fsSL https://packages.sury.org/php/README.txt | bash - \
&& apt-get install -y libc6 libc6-dev

WORKDIR /app

COPY --from=builder /app/bin/myapi .

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
  CMD curl --silent --fail http://localhost:8080/health || exit 1

CMD ["./myapi"]
