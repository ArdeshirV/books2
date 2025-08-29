FROM golang:1.25.0 AS builder
RUN apt-get update && apt-get install -y upx && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app .
RUN upx --best app  # RUN upx --ultra-brute app  # Best 4 production release

FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
# COPY --from=builder /app/assets ./assets  # If you have configs/assets to copy, do it here:
EXPOSE 64640
ENTRYPOINT ["./app"]
# docker build -t app .  # Build
# docker run --rm app    # Run

