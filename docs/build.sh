CUR=$(pwd)/$(dirname $0)

if test "$1" = 'clean'; then
    echo "Removing node_modules"
    rm -rf $CUR/docsjs/node_modules/
fi

rm -r $CUR/docsjs/dist
echo "Building docsjs"
cd $CUR/docsjs && yarn && yarn build
