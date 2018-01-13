FROM golang:1.8-alpine
ENV sourcesdir /go/src/github.com/learning-microservice/account/

COPY . ${sourcesdir}
RUN apk update
RUN apk add git
RUN go get -v github.com/golang/dep/cmd/dep && cd ${sourcesdir} && dep ensure && go install

ENTRYPOINT account
EXPOSE 18080
