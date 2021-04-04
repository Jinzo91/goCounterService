FROM golang:1.16

# Set the Current Working Directory inside the container
WORKDIR /Go

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .
COPY go.mod . 
COPY go.sum .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o restApi .

# Run the executable
CMD ["./restApi"]