FROM golang:stretch AS build
USER root
RUN apt-get update \
    && apt-get install -y \
    git
RUN git clone https://github.com/JesterOrNot/gosh.git \
    && cd gosh \
    && ./setup.sh
FROM debian
COPY --from=build gosh /usr/bin/
CMD gosh