FROM golang:1.20.3-alpine

# Set the working directory and move
WORKDIR /app

# Set the GO111MODULE environment variable
ENV GO111MODULE=on

# Copy go.mod and the app directory
COPY app/ ./

# Download dependencies and build the air tool
RUN go mod download && \
    go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

# Run the application using air
CMD ["air", "-c", ".air.toml"]
