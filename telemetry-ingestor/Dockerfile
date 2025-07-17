# Build stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o telemetry-ingestor ./cmd

# Final stage
FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/telemetry-ingestor .
CMD ["./telemetry-ingestor"]