API=/opt/homebrew/lib/node_modules/vuetify/dist/json/web-types.json
npm -g install vuetify@latest
find *.go | grep -v "fix-" | xargs rm
cat $API | vuetifyapi2go -comp=all
curl https://cdn.jsdelivr.net/npm/vuetify@3.x/dist/vuetify.min.js > dist/vuetify.min.js
curl https://cdn.jsdelivr.net/npm/vuetify@3.x/dist/vuetify.min.css > dist/vuetify.min.css
cd vuetifyjs && pnpm upgrade && pnpm build
