FROM golang:1.13.2-alpine3.10

WORKDIR /go/src/app

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# # Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]