FROM golang:1.26

WORKDIR /app

COPY go.mod ./

# RUN go mod download

COPY main.go ./
COPY static/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /dashboard_frontend

EXPOSE 5555

CMD ["/dashboard_frontend"]
