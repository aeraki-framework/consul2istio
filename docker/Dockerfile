FROM ubuntu:bionic
# Base image for debug builds.
# Built manually uploaded as "istionightly/base_debug"

# Do not add more stuff to this list that isn't small or critically useful.
# If you occasionally need something on the container do
# sudo apt-get update && apt-get whichever

# hadolint ignore=DL3005,DL3008

ENV consulAddress="127.0.0.1:8500"

RUN apt-get update && \
  apt-get install --no-install-recommends -y \
  ca-certificates \
  curl \
  iptables \
  iproute2 \
  iputils-ping \
  knot-dnsutils \
  netcat \
  tcpdump \
  conntrack \
  bsdmainutils \
  net-tools \
  lsof \
  linux-tools-generic \
  sudo \
  && update-ca-certificates \
  && apt-get upgrade -y \
  && apt-get clean \
  && rm -rf  /var/log/*log /var/lib/apt/lists/* /var/log/apt/* /var/lib/dpkg/*-old /var/cache/debconf/*-old

COPY consul2istio /usr/local/bin/
ENTRYPOINT /usr/local/bin/consul2istio -consulAddress=$consulAddress