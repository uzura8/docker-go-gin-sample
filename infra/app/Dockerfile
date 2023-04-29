FROM golang:1.20.3-alpine

# Make directory and move to directory
WORKDIR /go/app

#COPY app/go.mod .
#COPY app/go.sum .
COPY app/* ./

# Update and install git
#RUN apt-get update && apt-get install -y git

#RUN go mod download
RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
#CMD go run main.go
#CMD ["go", "run", "main.go"]