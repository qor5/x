ROOT=$(pwd)

PKGS="
$ROOT/../../web/corejs 
$ROOT/tiptap/tiptapjs 
$ROOT/vuetifyx/vuetifyxjs
$ROOT/cropper/cropperjs
$ROOT/redactor/redactorjs
"
for i in $PKGS
do
    echo "$i" && \
    cd $i && pnpm update && pnpm install && pnpm build
done
