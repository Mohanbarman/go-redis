FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o ./bin/redis-clone .

EXPOSE $PORT

CMD "bin/redis-clone" $PORT

