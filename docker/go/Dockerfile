FROM golang:alpine

COPY ./data-api /app
WORKDIR /app

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -o main

CMD ["./main"]