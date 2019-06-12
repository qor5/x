if test "$1" = 'clean'; then
    echo "Removing node_modules"
    rm -rf ./overlayjs/node_modules/
fi

CUR=$(pwd)
rm -r $CUR/overlayjs/dist
echo "Building overlayjs"
cd $CUR/overlayjs && yarn && yarn build
cd $CUR && packr
