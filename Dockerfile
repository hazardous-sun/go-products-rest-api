FROM docker.io/golang:1.23

# Set working directory
WORKDIR /go/src/app

# Copy the source code
COPY .. .

# Expose the port used for the requests
EXPOSE 8000

# Build the Go app
RUN go build -o main cmd/main.go

# Run the executable
CMD ["./main"]