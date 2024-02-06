# Use a multi-stage build to keep the final image minimal
FROM golang:1.21 AS builder

# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.* ./

# Download dependencies
RUN go mod download

# Copy the local package files to the container's working directory
COPY . .

# Build the binary (assuming main.go is in the cmd folder)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./app/cmd

# Use a minimal base image for the final stage
FROM alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/main .

# Copy config.json to the working directory
COPY ./config/config.json .

# Expose the port on which the application will run (adjust as needed)
EXPOSE 8090

ENV TZ=Asia/Bangkok
# Command to run the executable
CMD ["./main"]