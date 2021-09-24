API=/opt/homebrew/lib/node_modules/vuetify/dist/json/web-types.json
if [ -f "$API" ]; then
  echo ""
else
  npm -g install vuetify@2
fi
cat $API | vuetifyapi2go -comp=all
