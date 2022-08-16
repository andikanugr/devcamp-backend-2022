FROM golang:1.17-alpine

WORKDIR /app

COPY . ./

RUN go mod download

ENV ENV=Docker

RUN go build -o /devcamp-2022

EXPOSE 8080

CMD ["/devcamp-2022"]
