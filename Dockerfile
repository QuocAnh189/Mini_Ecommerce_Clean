# Step 1: Modules caching
FROM golang:1.24.1-alpine3.21 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.24.1-alpine3.21 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN mkdir -p /bin && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/app

# Step 3: Final
FROM alpine:latest
COPY --from=builder /app/policy /policy
COPY --from=builder /app/configs /configs
COPY --from=builder /bin/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Add curl to check MinIO
RUN apk add --no-cache curl

# Copy entrypoint.sh into container
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
