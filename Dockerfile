FROM golang:1.21-bookworm

MAINTAINER github@mullnet.org

COPY Dockerfile /Dockerfile

COPY . /setup/

RUN apt-get update \
    && apt-get -y upgrade \
    && cd /setup \
    && make clean \
    && make \
    && make install \
    && cd / \
    && rm -rf /setup

ENTRYPOINT ["/usr/local/bin/agent"]
