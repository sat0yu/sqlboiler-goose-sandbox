ARG GO_VERSION
FROM golang:${GO_VERSION}

ENV LANG=C.UTF-8 \
  LC_ALL=C.UTF-8 \
  GO111MODULE=on

RUN apt-get update && apt-get -y install postgresql
