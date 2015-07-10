FROM golang:latest

RUN mkdir -p /home/deployer/gosrc/src/github.com/mickelsonm/hn-test-app
ADD . /home/deployer/gosrc/src/github.com/mickelsonm/hn-test-app
WORKDIR /home/deployer/gosrc/src/github.com/mickelsonm/hn-test-app
RUN export GOPATH=/home/deployer/gosrc && go get
RUN export GOPATH=/home/deployer/gosrc && go build

ENTRYPOINT /home/deployer/gosrc/src/github.com/mickelsonm/hn-test-app/hn-test-app

EXPOSE 3000
