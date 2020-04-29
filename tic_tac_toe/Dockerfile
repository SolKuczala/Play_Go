FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && apt-get install -y golang git

RUN go get github.com/c-bata/go-prompt

COPY / /tic-tac-toe/
WORKDIR /tic-tac-toe/src/cli/

RUN go build *.go

ENTRYPOINT ./cli