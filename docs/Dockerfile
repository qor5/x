FROM golang:alpine as builder
RUN apk update && apk add git gcc libc-dev sqlite sqlite-dev && rm -rf /var/cache/apk/*
ARG GITHUB_TOKEN
WORKDIR /go/src/github.com/goplaid/x
COPY . .
RUN set -x && go get -d -v ./docs/docsmain/...
RUN GOOS=linux GOARCH=amd64 go build -o /app/entry ./docs/docsmain/

FROM alpine
RUN apk update && apk add sqlite sqlite-dev && rm -rf /var/cache/apk/*
COPY --from=builder /app/entry  /bin/docsmain
CMD /bin/docsmain
