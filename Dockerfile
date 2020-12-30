FROM golang:latest

RUN mkdir /go/src/boty
WORKDIR /go/src/boty

COPY go.mod go.sum ./
RUN go mod download

ADD . /go/src/boty

CMD ["go", "run", "."]