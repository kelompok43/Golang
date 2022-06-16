#builder
FROM golang:1.17 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o api main.go

FROM debian:11-slim

COPY --from=builder /app/api /app/
COPY --from=builder /app/config.env /app/

WORKDIR /app

EXPOSE 9600

CMD ["./api"]
