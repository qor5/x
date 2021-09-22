function appRestart() {
  echo "=================>"
  killall GoplaidPackageName
  cd front && npm run build && cd .. && \
  go build -o /tmp/GoplaidPackageName main.go && /tmp/GoplaidPackageName
}

export -f appRestart

find . -name *.go | entr -r bash -c "appRestart"

