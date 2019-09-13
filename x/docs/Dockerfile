FROM golang as builder
ARG GITHUB_TOKEN
WORKDIR /go/src/github.com/goplaid/
COPY . .
RUN set -x && \
    go get -d -v ./x/docs/docsmain/... && \
    GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o /app/entry ./x/docs/docsmain/

FROM alpine
COPY --from=builder /app/entry  /bin/docsmain
CMD /bin/docsmain
