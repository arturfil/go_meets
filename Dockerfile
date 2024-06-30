FROM golang:1.21.0-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o meetings_server ./cmd/server/main.go

EXPOSE 8080

CMD ["./meetings_server"]
