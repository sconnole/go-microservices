FROM golang:1.23-alpine AS builder

WORKDIR /app
RUN apk add --no-cache gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o job-creator

# Final minimal image
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache 
COPY --from=builder /app/job-creator /app/job-creator
EXPOSE 8080
CMD ["/app/job-creator"]