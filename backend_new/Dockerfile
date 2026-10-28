FROM golang:1.27 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY main.go go.mod go.sum ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /dashboard_backend


FROM alpine:3.24

WORKDIR /root/

COPY --from=builder ./dashboard_backend .

EXPOSE 8080

CMD ["./dashboard_backend"]
