# Build stage
FROM golang:1.23-bullseye AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o iotsimulator ./cmd/simulator

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/iotsimulator .
CMD ["./iotsimulator"]