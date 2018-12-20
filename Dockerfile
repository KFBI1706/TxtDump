FROM golang:latest
ENV GO111MODULE on
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app

RUN go get -v
