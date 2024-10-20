# Build stage
FROM golang:latest AS builder

WORKDIR /app

# Copy all files to the container
COPY . .

# Make sure build.sh is executable and run it
RUN chmod +x build.sh && ./build.sh

# Run stage
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/bin/api .

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["./api"]
