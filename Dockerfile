# Set GO Version
FROM golang:1.21.4

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]