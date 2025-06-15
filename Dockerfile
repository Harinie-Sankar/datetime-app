# Start with Go base image
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Copy code
COPY . .

# Build the Go app
RUN go build -o datetime-app .

# Expose the port
EXPOSE 8080

# Run the binary
CMD ["./datetime-app"]
