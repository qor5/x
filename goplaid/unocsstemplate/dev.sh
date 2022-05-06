function appRestart() {
  echo "=================>"
  killall GoplaidPackageName
  unocss **/*.go && \
  go build -o /tmp/GoplaidPackageName main.go && /tmp/GoplaidPackageName
}

export -f appRestart

find . -name *.go | entr -r bash -c "appRestart"

