FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && apt-get install -y golang git

RUN go get github.com/c-bata/go-prompt

COPY / /go/src/github.com/SolKuczala/Play_Go/tic_tac_toe/

WORKDIR /go/src/github.com/SolKuczala/Play_Go/tic_tac_toe/cli/

#RUN go build *.go
ENTRYPOINT bash
#ENTRYPOINT ./cli