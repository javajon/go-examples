FROM golang:1.24-alpine3.22
WORKDIR /app
COPY . .
RUN go mod init main
RUN go build -o main .
CMD ["./main"]
