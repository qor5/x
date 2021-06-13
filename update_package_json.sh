ROOT=$(pwd)

PKGS="
$ROOT/../web/corejs 
$ROOT/codehighlight/codehighlightjs 
$ROOT/overlay/overlayjs
$ROOT/tiptap/tiptapjs 
$ROOT/vuetify/vuetifyjs
$ROOT/docs/docsjs
"
for i in $PKGS
do
    echo "$i" && \
    rm $i/yarn.lock
    cd $i && ncu -u && yarn && \
    cd $i/.. && ./build.sh
done
