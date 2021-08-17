DIR=$(PWD)
cd $DIR/../ && \
snippetgo -pkg=examples > ./x/docs/examples/examples-generated.go
cd $DIR

function docsRestart() {
  echo "=================>"
  killall goplaidxdocs
#  export DEV_CORE_JS=1
#  export DEV_VUETIFY_JS=1
  go build -o /tmp/goplaidxdocs docs/docsmain/main.go && /tmp/goplaidxdocs
}

export -f docsRestart

find . -name *.go | entr -r bash -c "docsRestart"

