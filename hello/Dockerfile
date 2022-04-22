FROM golang:1.18.1-buster
WORKDIR /app
COPY . .
RUN go mod init main
RUN go build -o main .
CMD ["./main"]
