FROM golang

ADD . /go/src/github.com/briansimoni/chat

WORKDIR /go/src/github.com/briansimoni/chat

RUN go get

RUN go install github.com/briansimoni/chat

ENV PORT 4000

ENTRYPOINT /go/bin/chat

EXPOSE 4000