FROM golang:1.16.6-buster
WORKDIR /app
COPY . .
RUN go mod init main
RUN go build -o main .
CMD ["./main"]
