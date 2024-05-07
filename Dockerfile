# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local files into the container
COPY . .

# Build the application
RUN go build main.go

# Expose the ports required by the application
EXPOSE $PORT
EXPOSE $WS
EXPOSE $P2P

# Define the command to run when the container starts
# If REMOTE environment variable is set, run the application with remote connection parameters
# Otherwise, run the application with only the ID parameter
CMD ["sh", "-c", "if [ -n \"$REMOTE\" ]; then ./main --remote ${REMOTE} --id ${ID} --p2p ${P2P} --ws ${WS}; else ./main --id ${ID}; fi"]
