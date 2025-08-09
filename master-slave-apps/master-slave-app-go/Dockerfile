# Stage 1: building image using golang:1.24-alpine3.22
FROM golang:1.24-alpine3.22 AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY app/ .

# Download dependencies
RUN go mod tidy &&  go mod vendor && go mod download

# Build the application
RUN go build -v -o app ./cmd 

# Stage 2: Create the production image
FROM alpine:latest

# Install tzdata for timezone support
RUN apk add --no-cache tzdata

# Set the timezone environment variable (can be overridden by .env)
ENV TZ=Asia/Jakarta

# Configure the timezone
RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Set working directory
WORKDIR /app

# Create user and group for application
#
# Create a group with GID 1001
RUN addgroup -g 1001 binarygroup
# Create a user with UID 1001 and assign them to the 'binarygroup' group
RUN adduser -D -u 1001 -G binarygroup userapp

# Copy the binary from the builder stage
COPY --from=builder --chown=userapp:binarygroup /app/app .
COPY --from=builder /app/.env .

# Switch to the userapp user
USER userapp

# Command to run the application
CMD ["./app"]

