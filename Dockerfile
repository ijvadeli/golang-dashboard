# -- Compile the app
# Build Go app
FROM golang:1.27.2-alpine AS builder

WORKDIR /golang-dashboard

# Copy dependencies first
COPY go.mod ./
RUN go mod download

# Copy app source code
COPY . .

# Compile app
RUN CGO_ENABLED=0 GOOS=linux go build -o /dashboard .

# -- Contains only compiled program and minimal os
# Create final image
FROM alpine:3.22

WORKDIR /golang-dashboard

# Copy only compiled app
COPY --from=builder /dashboard .

# Port app uses
EXPOSE 8080

# Start API
CMD ["./dashboard"]
