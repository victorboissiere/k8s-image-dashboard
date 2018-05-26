FROM golang:1.10.2-alpine3.7

RUN apk --update add git
WORKDIR /opt/app
COPY . /opt/app

RUN go get -d .
RUN go build
RUN ls -la

FROM alpine:3.7

COPY --from=0 /opt/app /opt/app

ENV PORT 3000
RUN /opt/app/main
