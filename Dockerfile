FROM golang:1.13.12-alpine3.12

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/treeder/dockergo/

ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o myapp; cp myapp /app/

ENTRYPOINT ["./myapp"]