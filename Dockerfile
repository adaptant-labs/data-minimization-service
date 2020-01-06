FROM golang:latest as builder

ENV GO111MODULE=on

WORKDIR /go/src
ADD . /go/src

RUN go mod download
RUN go build -ldflags "-linkmode external -extldflags -static" -a -o /go/bin/app

FROM scratch
COPY --from=builder /go/bin/app /

EXPOSE 8080
ENTRYPOINT ["/app" ]
