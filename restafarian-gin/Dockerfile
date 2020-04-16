# Obtain certs, user and group for final stage
FROM alpine:3.11.5 as authority
RUN addgroup -r appuser && adduser -r -s /bin/false -g appuser appuser
RUN apk --no-cache add ca-certificates

# Build app binary for final stage
FROM golang:1.14.2-alpine3.11 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# Final stage
FROM scratch
COPY --from=authority /etc/group /etc/passwd /etc/
COPY --from=authority /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /main ./
USER appuser:appuser
EXPOSE 8080
ENTRYPOINT ["./main"]
