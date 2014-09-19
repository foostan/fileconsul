FROM ubuntu:trusty

MAINTAINER foostan ks@fstn.jp

RUN apt-get update
RUN apt-get install -y unzip wget curl

RUN mkdir /consul

RUN cd /consul && \
    wget https://dl.bintray.com/mitchellh/consul/0.4.0_linux_amd64.zip && \
    unzip 0.4.0_linux_amd64.zip && \
    mv consul /usr/local/bin

RUN cd /consul && \
    wget https://dl.bintray.com/mitchellh/consul/0.4.0_web_ui.zip && \
    unzip 0.4.0_web_ui.zip && \
    mv dist ui

ADD bin /consul/bin
ADD config /consul/config

CMD ["/bin/bash"]
