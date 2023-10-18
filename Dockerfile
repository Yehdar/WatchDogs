FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN mkdir /shared

VOLUME /shared

CMD ["sh", "-c", "go run server/main.go & go run client/main.go"]
