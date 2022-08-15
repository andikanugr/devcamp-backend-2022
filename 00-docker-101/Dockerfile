FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

ENV ENV=Docker

RUN go build -o /devcamp-2022

EXPOSE 8080

CMD ["/devcamp-2022"]
