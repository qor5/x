if test "$1" = 'clean'; then
    echo "Removing node_modules"
    rm -rf ./corejs/node_modules/
fi

CUR=$(pwd)
rm -r $CUR/corejs/dist
echo "Building corejs"
cd $CUR/corejs && yarn && yarn build
cd $CUR && packr
