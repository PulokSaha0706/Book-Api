# ---------- Stage 1: Build ----------
FROM golang:1.24 AS builder

WORKDIR /app

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o BookApi .

# ---------- Stage 2: Runtime ----------
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/BookApi .

EXPOSE 8000

ENTRYPOINT ["./BookApi"]
