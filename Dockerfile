FROM golang:alpine

ADD . /go/src/github.com/betandr/schwag-exchange

RUN go install github.com/betandr/schwag-exchange

ENTRYPOINT /go/bin/schwag-exchange

EXPOSE 8888
