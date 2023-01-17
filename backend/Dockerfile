FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go* ./
RUN go mod download

COPY . ./

RUN go build -v -o backend

FROM alpine:3

# Copy our static executable.
COPY --from=builder /app/backend /app/backend

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080


WORKDIR /app

# Run the Go Gin binary.
ENTRYPOINT ["/app/backend"]