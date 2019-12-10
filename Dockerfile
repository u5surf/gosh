FROM golang:stretch AS build
USER root
RUN git clone https://github.com/JesterOrNot/gosh.git \
    && cd gosh/src \
    && go get -v -t -d ./... \
    && go build -o /gosh *.go
FROM debian
COPY --from=build /gosh /usr/bin/
RUN mkdir \
        /go \
        /go/bin \
    && touch history.txt \
    && mv history.txt /go/bin
CMD [ "gosh" ]
