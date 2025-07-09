FROM golang:1.24.1-alpine AS builder

# Install git dan wire
RUN apk add --no-cache git
RUN apk add --no-cache curl

# Install wire
RUN go install github.com/google/wire/cmd/wire@latest

# Install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz \
    | tar -xz -C /usr/local/bin

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Run wire (auto-generate file)
RUN /go/bin/wire ./common/injector

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-app-server

RUN ls -all

#FROM scratch
FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/todo-app-server /
COPY --from=builder /app/config.yaml /app/config.yaml
COPY --from=builder /app/common/database/migrations /app/common/database/migrations

# Copy migrate binary
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

# Copy entrypoint
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENV DATABASE_URL="postgres://root:secret@postgres:5432/todo_app?sslmode=disable"

ENTRYPOINT ["/todo-app-server"]
