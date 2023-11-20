FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/ \
    go install github.com/go-jet/jet/v2/cmd/jet@latest

CMD sh deployments/startup.sh
