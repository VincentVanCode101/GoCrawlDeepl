# Use the official Golang base image
FROM golang:1.18 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY main.go main.go

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o run-app main.go

# Use a minimal alpine image for the final stage
FROM alpine:latest  

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/run-app .

# Command to run the executable
CMD ["./run-app"]
