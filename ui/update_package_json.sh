ROOT=$(pwd)

PKGS="
$ROOT/../../web/corejs 
$ROOT/codehighlight/codehighlightjs 
$ROOT/overlay/overlayjs
$ROOT/tiptap/tiptapjs 
"
for i in $PKGS
do
    echo "$i" && \
    rm $i/pnpm-lock.yaml
    cd $i && pnpm update && pnpm install && \
    cd $i/.. && ./build.sh
done
