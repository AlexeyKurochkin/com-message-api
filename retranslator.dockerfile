# Builder

ARG GITHUB_PATH=github.com/ozonmp/com-message-api

FROM golang:1.16-alpine AS builder
WORKDIR /home/${GITHUB_PATH}
RUN apk add --update make git protoc protobuf protobuf-dev curl
COPY Makefile Makefile
COPY . .
RUN make build-retranslator

# gRPC Server

FROM alpine:latest as server
LABEL org.opencontainers.image.source https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/retranslator .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .

RUN chown root:root retranslator

CMD ["./retranslator"]
