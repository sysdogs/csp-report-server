FROM golang:1.15.3-alpine3.12 AS build

COPY . /go/src/csp-report-server

WORKDIR /go/src/csp-report-server

RUN apk add --no-cache git

RUN go get && \
    go build

FROM alpine:3.12

COPY --from=build /go/src/csp-report-server/csp-report-server /csp-report-server

CMD ["/csp-report-server"]
