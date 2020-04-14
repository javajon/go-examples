# Obtain certs for final stage
FROM alpine:3.11.5 as certs
RUN apk --no-cache add ca-certificates

# Build app binary for final stage
FROM golang:1.14.2-alpine3.11 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# Final stage
FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /main ./
ENTRYPOINT ["./main"]
