FROM golang:1.15.5-alpine3.12

ENV CGO_ENABLE 0
ENV GO111MODULE on

ADD . /opt

WORKDIR /opt

RUN go env -w GOPROXY=https://goproxy.io,direct

RUN go build -o build/main github.com/crwkey/go-common

CMD ["./build/main"]
