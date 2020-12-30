FROM golang:latest

RUN mkdir /go/src/boty
WORKDIR /go/src/boty

COPY go.mod go.sum ./
RUN go mod download

ARG GOOGLE_APPLICATION_CREDENTIALS="boty-dev-firebase-adminsdk.json"
ENV GOOGLE_APPLICATION_CREDENTIALS=${GOOGLE_APPLICATION_CREDENTIALS}

ADD . /go/src/boty

CMD ["go", "run", "."]