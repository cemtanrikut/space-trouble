# Dockerfile

# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH environment variable at /go.
FROM golang:1.16 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -o spacetrouble specifies the output file name, replacing the default.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o spacetrouble ./cmd/main.go

# Use the official Debian slim image for a lean production container.
FROM debian:buster-slim
WORKDIR /app

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/spacetrouble /app/spacetrouble

# Run the web service on container startup.
CMD ["./spacetrouble"]
