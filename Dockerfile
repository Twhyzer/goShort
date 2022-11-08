## Build
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY ./src/cmd/main.go ./

RUN go build -o /build

EXPOSE 80

ENTRYPOINT ["/build"]