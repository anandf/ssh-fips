# --- Stage 1: Build the Go application ---
# We use a Go base image to compile the source code.
# golang:1.22-alpine is a good choice for a small build environment.
FROM registry.redhat.io/rhel8/go-toolset:1.24.4 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source code
COPY . .

# Compile the application.
# The `-o` flag specifies the output name of the binary.
RUN GOFIPS140=latest CGO_ENABLED=1 GOOS=linux GOEXPERIMENT=strictfipsruntime go build -tags strictfipsruntime -buildvcs=false -o /tmp/github-ssh-client .

# --- Stage 2: Create the final production image ---
# Use a minimal base image like 'alpine' to keep the image small.
#FROM registry.access.redhat.com/ubi8/ubi-minimal
FROM registry.redhat.io/rhel8/go-toolset:1.24.4

# Set the working directory for the final application
WORKDIR /root

# Copy the compiled binary from the builder stage
# We're copying from the "builder" stage to the current directory.
COPY --from=builder /tmp/github-ssh-client .

# Set the entrypoint to the compiled application
ENTRYPOINT ["./github-ssh-client"]
