FROM golang

MAINTAINER Husni Adil Makmur <husni.adil@gmail.com>

COPY src /go/src
COPY vendor/manifest /go/vendor/manifest
COPY Makefile /go/Makefile
COPY README.md /go/Dockerfile
COPY LICENSE /go/LICENSE
COPY README.md /go/README.md

WORKDIR /go

RUN make setup
RUN make build

ENTRYPOINT /go/bin/gsm

EXPOSE 8080
