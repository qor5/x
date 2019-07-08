if test "$1" = 'clean'; then
    echo "Removing node_modules"
    rm -rf ./corejs/node_modules/
fi

CUR=$(pwd)
rm -r $CUR/corejs/dist
echo "Building corejs"
cd $CUR/corejs && yarn && yarn build

echo "Building branoverlay"
cd $CUR/overlay/ && ./build.sh

echo "Building codehighlight"
cd $CUR/codehighlight/ && ./build.sh

echo "Building vuetify"
cd $CUR/vuetify/ && ./build.sh

cd $CUR && packr
