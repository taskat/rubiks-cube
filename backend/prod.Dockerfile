FROM golang:1.20 AS builder
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o ./bin/server ./src/main.go

FROM alpine:latest AS final
WORKDIR /app
COPY --from=builder /app/bin/server /bin/server

EXPOSE 8080
ENTRYPOINT ["/bin/server"]