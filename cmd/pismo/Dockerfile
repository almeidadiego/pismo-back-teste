FROM golang:1.15

ENV CGO_ENABLED=0

# Move to working directory
WORKDIR /go/src/pismo-back-teste

# Copy the code into the container
COPY . .

CMD ["go", "run", "./cmd/pismo/main.go"]