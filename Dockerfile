FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o authsvc

EXPOSE 8080

CMD ["./authsvc"]
