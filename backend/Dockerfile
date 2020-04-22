FROM golang:alpine
RUN apk add build-base
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
EXPOSE 3000

COPY . ${GOPATH}/src/github.com/ROWA/rowa-dev/backend
RUN cd ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/keys/ && update-ca-certificates

WORKDIR ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/src/main

CMD ["go", "run", "server.go"]
