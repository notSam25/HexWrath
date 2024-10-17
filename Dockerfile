# Use a Go base image with version 1.23 or higher
FROM golang:1.23-alpine

# Install required packages
RUN apk add --no-cache curl git

# Set the working directory inside the container
WORKDIR /HexWrath/src  # Update to the directory containing your Go files

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod ./
COPY go.sum ./
COPY src/templates/ ./
# Download dependencies
RUN go mod download

# Install air for hot-reloading
RUN go install github.com/air-verse/air@latest

# Copy the rest of the app's source code
COPY . .

# Expose the port the app runs on
EXPOSE 8080

# Start the app with Air hot-reloading
CMD ["air"]
