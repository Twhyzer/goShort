FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/src/cmd

RUN go build -o /build

EXPOSE 80



ENTRYPOINT ["/build"]