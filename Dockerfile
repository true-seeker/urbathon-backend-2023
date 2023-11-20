FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

#RUN go install github.com/golang-migrate/migrate/v4
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
#RUN go install github.com/go-jet/jet/v2/cmd/jet@latest

RUN ./migrate -source file://internal/app/storage/postgres/migrations -database postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable up
RUN jet -dsn="postgresql://$${POSTGRES_USER}:$${POSTGRES_PASSWORD}@$${POSTGRES_HOST}:$${POSTGRES_PORT}/$${POSTGRES_DB}?sslmode=disable" -schema=public -path=../../.gen

WORKDIR cmd/urbathon-backend-2023

RUN go build

