CUR=$(pwd)/$(dirname $0)

if test "$1" = 'clean'; then
    echo "Removing node_modules"
    rm -rf ./codehighlightjs/node_modules/
fi

rm -r $CUR/codehighlightjs/dist
echo "Building codehighlightjs"
cd $CUR/codehighlightjs && yarn && yarn build
