FROM ubuntu:latest
LABEL authors="Hope"

ENTRYPOINT ["top", "-b"]

# -------- Build Stage --------
FROM golang:1.21 AS builder

WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o main .

# -------- Final Stage --------
FROM gcr.io/distroless/base-debian11

WORKDIR /app

# Copy the built binary
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Command to run
CMD ["./main"]
