# Use Go 1.24.4 or latest
FROM golang:1.24.4-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o idgen ./cmd/main.go

EXPOSE 8080

CMD ["./idgen"]
