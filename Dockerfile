# Start with a base GoLang image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o test_axxonsoft

# Set the command to run when the container starts
CMD ["./test_axxonsoft"]