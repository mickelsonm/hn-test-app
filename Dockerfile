FROM google/golang

WORKDIR /gopath/src/github.com/mickelsonm/hn-test-app
ADD . /gopath/src/github.com/mickelsonm/hn-test-app

RUN go get
RUN go install

ENTRYPOINT ["/gopath/bin/hn-test-app"]

EXPOSE 3000
