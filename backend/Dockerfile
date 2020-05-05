FROM golang:alpine as build-stage
RUN apk add build-base
COPY . ${GOPATH}/src/github.com/ROWA/rowa-dev/backend
WORKDIR ${GOPATH}/src/github.com/ROWA/rowa-dev/backend/src/main
RUN go build server.go


FROM alpine as production-stage

#RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
EXPOSE 3000
RUN mkdir ROWA
COPY --from=build-stage /go/src/github.com/ROWA/rowa-dev/backend/src/main/ ./ROWA/main
COPY --from=build-stage /go/src/github.com/ROWA/rowa-dev/backend/src/settings/ ./ROWA/settings
RUN cd ROWA

WORKDIR ${HOME}/ROWA
ENTRYPOINT [ "./server" ]
