# Stage 1: Build the application
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy local files into the container
COPY . .

# Build the application
RUN go build main.go

# Stage 2: Create the final image
FROM debian:latest
# FROM debian:bullseye-slim

WORKDIR /app

# Copy the compiled application from the builder stage into the final image
COPY --from=builder /app/main .
RUN chmod +x main

# Expose the ports required by the application
EXPOSE $PORT
EXPOSE $WS
EXPOSE $P2P

# Define the command to run when the container starts
# If the REMOTE environment variable is set, run the application with remote connection parameters
# Otherwise, run the application with only the ID parameter
CMD ["sh", "-c", "if [ -n \"$REMOTE\" ]; then ./main --remote ${REMOTE} --id ${ID} --p2p ${P2P} --ws ${WS}; else ./main --id ${ID}; fi"]
