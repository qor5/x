CUR=$(pwd)/$(dirname $0)

if test "$1" = 'clean'; then
    echo "Removing node_modules"
    rm -rf ./vuetifyxjs/node_modules/
fi

rm -r $CUR/vuetifyxjs/dist
echo "Building vuetifyjs"
cd $CUR/vuetifyxjs && pnpm install && pnpm format && pnpm run build
