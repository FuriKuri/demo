FROM golang:1.10
ADD . /go/src/github.com/furikuri/demo
WORKDIR /go/src/github.com/furikuri/demo
RUN go get ./
RUN go build

FROM alpine:3.7
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /root/
COPY --from=0 /go/bin/demo .
COPY index.html .

ENV HTML_TITLE="Hello World"
ENV HTML_BODY="This is a sample Page."

ENTRYPOINT ["/root/demo"]