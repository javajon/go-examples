# Obtain certs for final stage
FROM alpine:3.22 AS authority
RUN mkdir /user && \
    echo 'appuser:x:1000:1000:appuser:/:' > /user/passwd && \
    echo 'appgroup:x:1000:' > /user/group
RUN apk --no-cache add ca-certificates

# Build app binary for final stage
FROM golang:1.24-alpine3.22 AS builder
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with limited parallelism and optimizations in lab environment
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  GOMAXPROCS=2

# Build the binary
RUN go build -ldflags="-w -s -extldflags '-static'" -tags netgo -a -o /main .
  
# Final stage
FROM scratch
COPY --from=authority /user/group /user/passwd /etc/
COPY --from=authority /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /main ./
USER appuser:appgroup
EXPOSE 8080
ENTRYPOINT ["./main"]
