# Obtain certs for final stage
FROM alpine:3.22 AS authority
RUN mkdir /user && \
    echo 'appuser:x:1000:1000:appuser:/:' > /user/passwd && \
    echo 'appgroup:x:1000:' > /user/group
RUN apk --no-cache add ca-certificates

# Build app binary for final stage
FROM golang:1.24-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# Final stage
FROM scratch
COPY --from=authority /user/group /user/passwd /etc/
COPY --from=authority /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /main ./
USER appuser:appgroup
EXPOSE 8080
ENTRYPOINT ["./main"]
