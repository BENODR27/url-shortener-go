FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download && go build -o url-shortener-go ./cmd/server

FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=builder /app/url-shortener-go .
EXPOSE 8080
CMD ["./url-shortener-go"]
