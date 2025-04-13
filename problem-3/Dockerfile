# Use the official Golang base image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy go app
COPY main.go .

# Build Go app
RUN go build -o app .

# Run the executable
CMD ["./app"]
