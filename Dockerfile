# use official Golang image
FROM golang:1.21

# set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o 401k_calculator .

RUN chmod +x /app/401k_calculator

#EXPOSE the port
EXPOSE 8080

# Run the executable
CMD ["/app/401k_calculator"]