# Go base image
FROM golang:1.21 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY src/ .

# Build the Go app
RUN go mod tidy
RUN go build -o pars ./pars.go

# Start a new stage from busybox
FROM busybox:latest

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/pars /usr/bin/pars

# Set /home as the working directory
WORKDIR /home

# Keep the container running
ENTRYPOINT ["sleep", "infinity"]