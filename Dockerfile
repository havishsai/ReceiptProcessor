# Use the official Golang image as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app from the specified directory
RUN go build -o receiptProcessor ./cmd/receiptProcessor/

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./receiptProcessor"]
