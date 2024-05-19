FROM golang:1.22.1-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/Waramoto/hryvnia-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/hryvnia-svc /go/src/github.com/Waramoto/hryvnia-svc


FROM alpine:3.19.1

COPY --from=buildbase /usr/local/bin/hryvnia-svc /usr/local/bin/hryvnia-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["hryvnia-svc"]
