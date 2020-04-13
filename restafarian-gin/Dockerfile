# Obtain certs for final stage
FROM alpine:3.11.5 as certs
RUN apk --no-cache add ca-certificates

# Build app binary for final stage
FROM golang:1.14.2-buster AS builder
WORKDIR $GOPATH/src/app
RUN go get -u github.com/gin-gonic/gin
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# Final stage
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /main ./
EXPOSE 8080
ENTRYPOINT ["./main"]
