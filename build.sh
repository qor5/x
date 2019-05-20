if test "$1" = 'clean'; then
    echo "Removing node_modules"
    rm -rf ./codehighlightjs/node_modules/
fi

CUR=$(pwd)
rm -r $CUR/codehighlightjs/dist
echo "Building codehighlightjs"
cd $CUR/codehighlightjs && yarn && yarn build
cd $CUR && packr
