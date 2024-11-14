# Golang 1.23 Image
FROM golang:1.23 AS dev

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD [ "air" ]


