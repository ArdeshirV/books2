# Build stage
FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go build -ldflags "-w -s" -o app .
#RUN pkg install -y upx
#RUN upx app

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]

