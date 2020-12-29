FROM golang:latest

RUN mkdir /go/src/boty
WORKDIR /go/src/boty
ADD . /go/src/boty

CMD ["go", "run", "."]