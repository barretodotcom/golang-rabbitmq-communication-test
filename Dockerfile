FROM golang:latest

WORKDIR /go/src/

COPY . /go/src/

ENTRYPOINT ["tail", "-f", "/dev/null"]