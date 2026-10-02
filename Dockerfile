FROM golang:1.27.1-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/exe ./cmd/app 

FROM alpine:3.24.2
WORKDIR /app
COPY --from=builder /app/exe /app/exe
CMD ["/app/exe"]
