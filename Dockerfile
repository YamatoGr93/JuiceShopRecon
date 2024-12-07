# Use the official Go image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project directory into the container
COPY . .

# Build the unified tool (main.go)
RUN go build -o recon_tool main.go

# Expose port 3000 for web testing (optional, if needed by Juice Shop)
EXPOSE 3000

# Command to run the tool
ENTRYPOINT ["./recon_tool"]
