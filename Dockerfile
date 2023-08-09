
FROM golang:1.21-alpine3.18 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./justcode cmd/justcode/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/justcode .

COPY --from=builder /app/api/. ./api

EXPOSE 8080

CMD ["./justcode"]
