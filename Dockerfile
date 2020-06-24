FROM ubuntu:latest
FROM golang:latest

ENV DEBIAN_FRONTEND=noninteractive
ENV GOPATH=/go
ENV MODE=api

RUN apt-get update -y && apt-get install -y git 

RUN go get github.com/c-bata/go-prompt github.com/gin-gonic/gin

COPY / ${GOPATH}/src/github.com/SolKuczala/tic-tac-go/

WORKDIR ${GOPATH}/src/github.com/SolKuczala/tic-tac-go/cli/
RUN go install

WORKDIR ${GOPATH}/src/github.com/SolKuczala/tic-tac-go/api/
RUN go install

WORKDIR ${GOPATH}/bin/

CMD ["/bin/bash", "true"]
