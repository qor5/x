CUR=$(pwd)/$(dirname $0)

mkdir -p $CUR/dist && curl -fsSL https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css > $CUR/dist/material-components-web.min.css

packr

