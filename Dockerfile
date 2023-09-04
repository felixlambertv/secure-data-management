FROM golang:1.20-alpine
WORKDIR /app

ENV GOOS="linux"
ENV CGO_ENABLED=0
ENV GO111MODULE="on"

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    git \
    && update-ca-certificates

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/vektra/mockery/v2@latest

EXPOSE 8080

# Run the air command in the directory where our code will live
ENTRYPOINT ["air", "-c", ".air.toml"]