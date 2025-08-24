# Build stage
FROM golang:1.25 as builder
WORKDIR /app
COPY . .
RUN go build -o app .

# Run stage
FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]

