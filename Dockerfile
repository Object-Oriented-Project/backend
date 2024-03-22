# write me a dockerfile for backend golang
FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 3001

CMD ["./main"]