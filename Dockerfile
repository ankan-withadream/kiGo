

# Use the golang:alpine image as the base
FROM golang:alpine

# Set the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Run go get to download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Set the Current Working Directory inside the container
#WORKDIR /

# Copy the source from the current directory to the Working Directory inside the container
#COPY . .


# RUN go mod init github.com/ankan-withadream/kiGo

# RUN go get .


# Build the Go app
RUN go build

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./kiGo"]





