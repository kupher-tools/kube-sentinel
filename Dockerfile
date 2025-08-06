FROM golang:1.24 AS builder
WORKDIR /app
COPY . . 
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/

FROM alpine:3.18
RUN adduser -D kupher
COPY --from=builder /app/main /usr/local/bin/main
USER kupher
EXPOSE 443
ENTRYPOINT ["/usr/local/bin/main"]