FROM golang:stretch AS build
USER root
RUN git clone https://github.com/JesterOrNot/gosh.git \
    && cd gosh/src \
    && go get -v -t -d ./... \
    && go build -o /gosh *.go
FROM scratch
COPY --from=build /gosh /usr/bin/
RUN mkdir /.gosh \
    && touch history.txt \
    && mv history.txt /.gosh
CMD [ "gosh" ]
