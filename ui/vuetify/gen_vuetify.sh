API=/opt/homebrew/lib/node_modules/vuetify/dist/json/web-types.json
npm -g install vuetify@3.6.14
find *.go | grep -v "fix-" | xargs rm
cat $API | vuetifyapi2go -comp=all
