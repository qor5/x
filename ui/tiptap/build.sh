# @snippet_begin(TiptapBuilderSH)
CUR=$(pwd)/$(dirname $0)

if test "$1" = 'clean'; then
    echo "Removing node_modules"
    rm -rf $CUR/tiptapjs/node_modules/
fi

rm -r $CUR/tiptapjs/dist
echo "Building tiptapjs"
cd $CUR/tiptapjs && npm install && npm run build

# @snippet_end
