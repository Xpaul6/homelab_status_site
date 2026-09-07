FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY main.go go.mod go.sum ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /dashboard_frontend


FROM alpine:latest

WORKDIR /root/

COPY --from=builder ./dashboard_frontend .

COPY static/ ./static/

EXPOSE 5555

CMD ["./dashboard_frontend"]
