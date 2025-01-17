# Base image
FROM golang:1.22

# Set the working directory
WORKDIR /app

# Copy source code
COPY . .

# Download dependencies and build the application
RUN go mod download
RUN go build -o job_portal main.go

# Expose application ports
EXPOSE 8080

# Start the job portal service
CMD ["./job_portal"]
