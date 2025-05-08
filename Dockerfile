# ---------- Stage 1: Build Go App ----------
FROM golang:1.22.5 AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build -o main .

# ---------- Stage 2: Prepare shell from BusyBox ----------
FROM busybox:latest AS busybox
# Just using this image to extract /bin/sh

# ---------- Stage 3: Distroless Runtime ----------
FROM gcr.io/distroless/base

WORKDIR /

# Copy Go binary from builder
COPY --from=builder /app/main .

# Copy static assets (optional)
COPY --from=builder /app/himmu ./static

# Copy /bin/sh from BusyBox (debugging purposes only!)
COPY --from=busybox /bin/sh /bin/sh

EXPOSE 8080

# Use Go binary as entrypoint
CMD ["./main"]

