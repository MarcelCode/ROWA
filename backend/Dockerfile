FROM golang:alpine as build-stage
RUN apk add build-base
COPY . ${GOPATH}/src/github.com/ROWA/rowa-dev/backend
WORKDIR ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/src/main
#CMD ["go", "build","-o ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/src/main", "server.go"]
RUN go build server.go


FROM golang:alpine as production-stage

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
EXPOSE 3000
COPY --from=build-stage ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/src/main ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/src/main
COPY --from=build-stage ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/keys ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/keys
#COPY . ${GOPATH}/src/github.com/ROWA/rowa-dev/backend
RUN cd ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/keys/ && update-ca-certificates

WORKDIR ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/src/main
ENTRYPOINT [ "./server" ]
