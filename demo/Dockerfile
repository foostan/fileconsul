FROM ubuntu:trusty

MAINTAINER foostan ks@fstn.jp

RUN apt-get update
RUN apt-get install -y unzip wget curl

RUN wget https://dl.bintray.com/mitchellh/consul/0.3.1_linux_amd64.zip
RUN unzip 0.3.1_linux_amd64.zip
RUN mv consul /bin/consul
