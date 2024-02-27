# Use Golang as base image
FROM golang:1.16-alpine as build

# Set working directory
WORKDIR /app

# Copy go.mod
COPY go.mod ./

# Download Go dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main ./src

# Alpine for a lightweight image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy executable from build stage
COPY --from=build /app/main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]