FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN apk add --no-cache git mercurial \
&& go install -v ./... \
&& apk del git mercurial

CMD ["app"]