FROM golang:1.14.2-buster
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]
