FROM golang:1.10.2-alpine3.7

RUN apk --update add git
WORKDIR /gopath/src/app
RUN go get github.com/justincampbell/timeago && \
    go get k8s.io/api/core/v1 && \
    go get k8s.io/client-go/rest && \
    go get k8s.io/client-go/tools/clientcmd && \
    go get k8s.io/apimachinery/pkg/apis/meta/v1 && \
    go get github.com/labstack/echo

COPY . /gopath/src/app/

RUN go build -ldflags "-s -w" -o /usr/bin/imageDashboard

FROM alpine:3.7

COPY --from=0 /usr/bin/imageDashboard /usr/bin/

ENV PORT 3000
ENTRYPOINT ["/usr/bin/imageDashboard"]
