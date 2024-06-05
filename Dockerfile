# Stage 1: Build
FROM golang:1.17 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Stage 2: Package
FROM gcr.io/distroless/static:1.5.2
WORKDIR /app
COPY --from=builder /app/main /app/main
CMD ["./main"]

# Expose port 8080
EXPOSE 8080

# Set entrypoint
ENTRYPOINT ["/app/main"]

# Set healthcheck
HEALTHCHECK --interval=10s --timeout=5s CMD curl -f http://localhost:8080/health || exit 1




