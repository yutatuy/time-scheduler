FROM golang:1.19.5-alpine

WORKDIR /app/go
COPY ./go-app .
RUN apk add --no-cache git \
    && go get -d github.com/cosmtrek/air \
    && cd $GOPATH/pkg/mod/github.com/cosmtrek/air* \
    && go build -o /go/bin/air \
    && rm -rf $GOPATH/pkg/mod/* 

CMD air -c .air.toml