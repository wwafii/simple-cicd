# simple-cicd/Dockerfile
FROM golang:1.23-alpine AS builder 
WORKDIR /app

# Copy deps dulu (layer caching)
COPY go.mod go.sum* ./
RUN go mod download

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-cicd .

# Runtime image (super kecil ~15MB)
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /simple-cicd .
EXPOSE 8080 
CMD ["./simple-cicd"]
