FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o proxy main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/proxy .
EXPOSE 8080
ENV PORT=8080
CMD ["./proxy"]
