FROM golang:1.22 AS build

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY go ./go
COPY main.go .

# Build the Go application
ENV CGO_ENABLED=0
RUN go build -o swagger .

# Create a minimal runtime image
FROM scratch AS runtime

WORKDIR /
COPY --from=build /app/swagger ./
EXPOSE 8080
ENTRYPOINT ["./swagger"]
