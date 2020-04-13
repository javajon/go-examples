FROM golang:1.14.2-buster
RUN mkdir /app
WORKDIR /app
COPY *.go /app/
RUN go build -o main .
CMD ["/app/main"]
