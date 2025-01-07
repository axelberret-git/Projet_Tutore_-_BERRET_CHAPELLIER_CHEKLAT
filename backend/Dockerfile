FROM golang:1.20.5 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
