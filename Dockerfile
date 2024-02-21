FROM golang:1.20.3-alpine AS builder

COPY . /github.com/dmtrybogdanov/chat-server/source
WORKDIR /github.com/dmtrybogdanov/chat-server/source

RUN go mod download
RUN go build -o ./bin/chat_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/dmtrybogdanov/chat-server/source/bin/ .

CMD ["./chat_server"]
