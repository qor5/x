CUR=$(pwd)/$(dirname $0)

if test "$1" = 'clean'; then
    echo "Removing node_modules"
    rm -rf ./overlayjs/node_modules/
fi

rm -r $CUR/overlayjs/dist
echo "Building overlayjs"
cd $CUR/overlayjs && yarn && yarn build
cd $CUR && packr
