#!/bin/sh

# Set the script to exit on error
set -e

# Print commands before executing them
set -x

echo "Building..."

# Ensure Go modules are tidy
go mod tidy

# Create the bin directory if it doesn't exist
mkdir -p bin

# Build the binary and check for errors
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/api web/main/main.go

echo "Build completed successfully!"
