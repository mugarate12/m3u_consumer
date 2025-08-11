# FROM golang:1.23-alpine
FROM golang:1.23

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

ENTRYPOINT ["go", "test", "./..."]
# ENTRYPOINT [ "bash" ]
