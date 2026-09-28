# Build stage
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o cloud-platform-api .

# Runtime stage
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/cloud-platform-api .

EXPOSE 8080

USER 10001

CMD ["./cloud-platform-api"]
