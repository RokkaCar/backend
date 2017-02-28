FROM golang

ADD . /go/src/github.com/RokkaCar/backend

RUN go install github.com/RokkaCar/backend

ENTRYPOINT ["/go/bin/backend"]

EXPOSE 8080

