FROM golang:1.15

RUN mkdir -p $GOPATH/src/github.com/adivinho/pmm-managed/tests

WORKDIR $GOPATH/src/github.com/adivinho/pmm-managed/tests/
COPY . $GOPATH/src/github.com/adivinho/pmm-managed/tests/

CMD make init run-race
