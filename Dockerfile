# Use Golang official image as base
FROM golang:1.20

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o clipboard-server .

# Expose port 8080
EXPOSE 8080

# Run the Go application
CMD ["./clipboard-server"]
