# Use the official Go image as the base
FROM golang:1.23

# Set working directory
WORKDIR /app

# Copy application files
COPY . .

# Build the Go application
RUN go build -o server main.go

# Expose the application port
EXPOSE 8080

# Start the server
CMD ["./server"]

