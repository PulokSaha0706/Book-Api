FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o BookApi .

EXPOSE 8000

ENTRYPOINT ["./BookApi"]
