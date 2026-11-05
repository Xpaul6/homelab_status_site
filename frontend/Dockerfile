FROM golang:1.27 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /dashboard_frontend


FROM alpine:3.24

WORKDIR /root/

COPY --from=builder ./dashboard_frontend .

COPY static/ ./static/

EXPOSE 5555

CMD ["./dashboard_frontend"]
