FROM golang:1.22-alpine AS builder

COPY . /github.com/bogdanove/chat-server/source/
WORKDIR /github.com/bogdanove/chat-server/source/

RUN go mod download
RUN go build -o ./bin/chat-server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/bogdanove/chat-server/source/bin/chat-server .

CMD ["./chat-server"]