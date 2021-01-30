FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src

COPY . /src
RUN cd /src
RUN go mod download

RUN go build -o bin/nubank main.go

ENTRYPOINT [  "./bin/nubank" ]
