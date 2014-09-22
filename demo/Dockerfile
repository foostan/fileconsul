FROM ubuntu:trusty

MAINTAINER foostan ks@fstn.jp

RUN apt-get update
RUN apt-get install -y unzip wget curl
RUN apt-get install -y apache2 ntp # install sample packages

RUN mkdir /consul

RUN cd /consul && \
    wget https://dl.bintray.com/mitchellh/consul/0.4.0_linux_amd64.zip && \
    unzip 0.4.0_linux_amd64.zip && \
    mv consul /usr/local/bin

RUN cd /consul && \
    wget https://dl.bintray.com/mitchellh/consul/0.4.0_web_ui.zip && \
    unzip 0.4.0_web_ui.zip && \
    mv dist ui

RUN cd /consul && \
    wget https://dl.bintray.com/mitchellh/consul/0.4.0_linux_amd64.zip && \
    wget http://dl.bintray.com/foostan/fileconsul/0.1.1_linux_amd64.zip && \
    unzip 0.1.1_linux_amd64.zip && \
    mv fileconsul /usr/local/bin

ADD share /consul/share

CMD ["/bin/bash"]
