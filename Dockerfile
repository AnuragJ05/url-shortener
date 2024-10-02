# Use the official Golang image
FROM golang:1.22

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the application
COPY . .
RUN go build -o url-shortener .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./url-shortener"]
