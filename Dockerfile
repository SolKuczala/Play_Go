FROM alpine:latest
FROM golang:latest

ENV DEBIAN_FRONTEND=noninteractive
ENV GOPATH=/go
ENV MODE=api

RUN apk update -y && apk add -y git

RUN go get github.com/c-bata/go-prompt github.com/gin-gonic/gin

COPY / ${GOPATH}/src/github.com/SolKuczala/tic-tac-go/

WORKDIR ${GOPATH}/src/github.com/SolKuczala/tic-tac-go/cli/
RUN go install

WORKDIR ${GOPATH}/src/github.com/SolKuczala/tic-tac-go/api/
RUN go install

WORKDIR ${GOPATH}/bin/

ENTRYPOINT ./${MODE}
