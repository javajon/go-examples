FROM golang:1.22-alpine3.19
WORKDIR /app
COPY . .
RUN go mod init main
RUN go build -o main .
CMD ["./main"]
