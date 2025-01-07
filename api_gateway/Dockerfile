FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make generate

RUN make build

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/api_gateway .

EXPOSE 8080

CMD ["./api_gateway"]