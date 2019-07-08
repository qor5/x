CUR=$(pwd)/$(dirname $0)

if test "$1" = 'clean'; then
    echo "Removing node_modules"
    rm -rf ./vuetifyjs/node_modules/
fi

rm -r $CUR/vuetifyjs/dist
echo "Building vuetifyjs"
cd $CUR/vuetifyjs && yarn && yarn build
cd $CUR && packr
