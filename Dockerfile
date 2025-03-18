# Set build-time arguments
ARG GO_VERSION=1.24.1
ARG ALPINE_VERSION=3.21

# Base stage with specific Alpine version and build args
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS base
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && \
    go mod verify


# Builder stage
FROM base AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-w -s" \
    -trimpath \
    -o app \
    main.go

# Runtime stage
FROM alpine:${ALPINE_VERSION} AS runtime

# Install runtime dependencies
RUN apk add --no-cache \
    ca-certificates \
    tzdata \
    && update-ca-certificates

# Create non-root user
RUN addgroup -S appgroup \
    && adduser -S appuser -G appgroup

# Switch to non-root user
USER appuser

# Set working directory
WORKDIR /home/appuser

# Copy migration files and binary
COPY --from=builder /src/app ./
COPY ./db/postgres/migration ./db/postgres/migration

# Run the application
CMD ["./app"]