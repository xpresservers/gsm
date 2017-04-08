FROM golang

MAINTAINER Husni Adil Makmur <husni.adil@gmail.com>

COPY src /go/src
COPY vendor/manifest /go/vendor/manifest
COPY Makefile /go/Makefile
WORKDIR /go
RUN make setup
RUN make build

ENTRYPOINT /go/bin/gsm

EXPOSE 8080
