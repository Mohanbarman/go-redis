FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o ./bin/go-redis .

EXPOSE $PORT

CMD "bin/go-redis" $PORT

