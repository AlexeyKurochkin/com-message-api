# Builder

ARG GITHUB_PATH=github.com/ozonmp/com-message-api

FROM golang:1.16-alpine AS builder
WORKDIR /home/${GITHUB_PATH}
RUN apk add --update build-base make git protoc protobuf protobuf-dev curl
COPY Makefile Makefile
COPY . .
RUN make build-omp-bot

RUN go get github.com/go-delve/delve/cmd/dlv

# omp-bot

FROM alpine:latest as server
LABEL org.opencontainers.image.source https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/omp-bot .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .

COPY --from=builder /go/bin/dlv .

RUN chown root:root omp-bot

EXPOSE 40001

#CMD ["./omp-bot"]
CMD ["./dlv", "--listen=:40001", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./omp-bot"]

