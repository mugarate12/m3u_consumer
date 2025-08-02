FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

ENTRYPOINT ["go", "run", "./app.go"]
