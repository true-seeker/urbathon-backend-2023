FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go install github.com/go-jet/jet/v2/cmd/jet@latest

RUN go mod download

COPY . .

WORKDIR cmd/urbathon-backend-2023

RUN go build