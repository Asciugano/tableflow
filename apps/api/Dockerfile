FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN go build -o /tableflow-api ./cmd/api

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /tableflow-api .

EXPOSE 8080

CMD ["./tableflow-api"]
