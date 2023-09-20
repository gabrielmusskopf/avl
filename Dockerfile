FROM golang:1.21 as builder

WORKDIR /app

# Copy Go module files
COPY go.* ./

# Download dependencies
RUN go mod download

# Copy source files
COPY ./cmd ./cmd
COPY ./http ./http
COPY *.go .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/avl ./cmd/*.go

FROM alpine:3.14.10

EXPOSE 3333

# Copy files from builder stage
COPY --from=builder /app/bin/avl .
COPY --from=builder /app/ .

# Run binary
ENTRYPOINT ["/avl"]
