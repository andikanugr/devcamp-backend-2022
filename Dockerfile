FROM golang:1.17.0-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

RUN export GO111MODULE=on

COPY . .

RUN go mod vendor

# Build the application
RUN go build -o main .

# Expose port 9000
EXPOSE 9000

# Run the executable
CMD ["./main"]