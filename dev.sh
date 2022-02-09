DIR=$(PWD)
# go install github.com/sunfmin/snippetgo@latest
cd $DIR/../ && \
snippetgo -pkg=examples > ./x/docs/examples/examples-generated.go
cd $DIR

function docsRestart() {
  echo "=================>"
  killall goplaidxdocs
#  export DEV_CORE_JS=1
#  export DEV_VUETIFY_JS=1
#  export DEV_PRESETS=1
  go build -o /tmp/goplaidxdocs docs/docsmain/main.go && /tmp/goplaidxdocs
}

export -f docsRestart
ulimit -n 1000000
find . -name "*.go" | entr -r bash -c "docsRestart"
