FROM golang:1.25.0 AS builder

# Install UPX
RUN apt-get update && \
    apt-get install -y upx && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Add go mod files and download dependencies early for Docker cache
COPY go.mod go.sum ./
RUN go get 

# Copy the rest of your app
COPY . .

# Build the binary (statically linked for scratch image)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o docker-test .

# Compress the binary with UPX
RUN upx --best --lzma docker-test



# --- Final stage (minimal image) ---
FROM scratch

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/docker-test .

# If you have configs/assets to copy, do it here:
# COPY --from=builder /app/assets ./assets

EXPOSE 64640

# Run the binary
ENTRYPOINT ["./docker-test"]


## To Build and Run:
# docker build -t my-golang-upx-app .
# docker run --rm my-golang-upx-app

