FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive
ENV GOPATH=/go
ENV HOST=127.0.0.1
ENV PORT=9090
ENV STRATEGY=random
ENV PLAYER=X
ENV PLAYFIRST=true

RUN apt-get update -y && apt-get install -y golang git

COPY / ${GOPATH}/src/github.com/SolKuczala/tic-tac-go/bot_player/

WORKDIR ${GOPATH}/src/github.com/SolKuczala/tic-tac-go/bot_player/
RUN go install

WORKDIR ${GOPATH}/bin/

ENTRYPOINT ./bot_player -ip=${HOST} -port=${PORT} -player=${PLAYER} -strategy=${STRATEGY} -playfirst=${PLAYFIRST}