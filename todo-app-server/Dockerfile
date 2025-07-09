FROM golang:1.24.1-alpine AS builder

# Install git dan wire
RUN #apk add --no-cache git

# Install wire
RUN #go install github.com/google/wire/cmd/wire@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Run wire (auto-generate file)
RUN #/go/bin/wire ./common/injector

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-app-server

#FROM scratch
FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/todo-app-server /
COPY --from=builder /app/config.yaml /app/config.yaml

ENTRYPOINT ["/todo-app-server"]
