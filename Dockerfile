FROM alpine:latest

RUN mkdir /app
WORKDIR /app
ADD teampay-greet-service /app/teampay-greet-service

CMD ["./teampay-greet-service"]