FROM golang:1.23-alpine AS builder
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -ldflags="-s -w" -o /app/answer ./cmd/answer

FROM alpine:3.19
COPY --from=builder /app/demoapp /app/demoapp
COPY --from=builder /build/internal/apps/demoapp/config/config.prod.yaml /app/config.yaml

WORKDIR /app
EXPOSE 8099

ENTRYPOINT ["/app/demoapp"]
