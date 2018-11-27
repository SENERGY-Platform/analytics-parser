FROM golang

RUN go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/analytics-parser
WORKDIR /go/src/analytics-parser

RUN dep ensure
RUN go build

EXPOSE 5001

CMD ./analytics-parser
