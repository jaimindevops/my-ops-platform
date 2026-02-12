# Stage 1: Build the Go App (No changes here)
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
# build specifically for the distroless architecture
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Run the App (SECURE CHANGE)
# Use 'distroless/static' instead of 'alpine'
# This removes the OS layer, shell, and package manager = No CVEs!
FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /app/main .

# Use a non-root user (Standard Security Practice)
USER 65532:65532

EXPOSE 8080

ENTRYPOINT ["/main"]
