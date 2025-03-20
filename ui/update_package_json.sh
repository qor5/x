ROOT=$(pwd)

PKGS="
$ROOT/../../web/corejs 
$ROOT/codehighlight/codehighlightjs 
$ROOT/tiptap/tiptapjs 
$ROOT/vuetifyx/vuetifyxjs
$ROOT/cropper/cropperjs
$ROOT/redactor/redactorjs
"
for i in $PKGS
do
    echo "$i" && \
    rm $i/pnpm-lock.yaml
    cd $i && pnpm update && pnpm install && pnpm build
done
