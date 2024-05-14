module github.com/qor5/x/v3

go 1.22.2

require (
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/iancoleman/strcase v0.3.0
	github.com/jinzhu/inflection v1.0.0
	github.com/lib/pq v1.10.9
	github.com/manifoldco/promptui v0.9.0
	github.com/markbates/goth v1.79.0
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/ory/ladon v1.3.0
	github.com/pquerna/otp v1.4.0
	github.com/qor5/web/v3 v3.0.0
	github.com/spf13/cast v1.6.0
	github.com/stretchr/testify v1.9.0
	github.com/sunfmin/reflectutils v1.0.3
	github.com/theplant/htmlgo v1.0.3
	github.com/theplant/osenv v0.0.1
	github.com/theplant/testingutils v0.0.0-20240326065615-ab2586803ce4
	golang.org/x/crypto v0.22.0
	golang.org/x/text v0.14.0
	gorm.io/driver/postgres v1.5.7
	gorm.io/driver/sqlite v1.5.5
	gorm.io/gorm v1.25.9
)

require (
	cloud.google.com/go/compute v1.25.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/boombuler/barcode v1.0.1 // indirect
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dlclark/regexp2 v1.11.0 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/gorilla/context v1.1.2 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/gorilla/sessions v1.2.2 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20231201235250-de7065d80cb9 // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/mrjones/oauth v0.0.0-20190623134757-126b35219450 // indirect
	github.com/ory/pagination v0.0.1 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/oauth2 v0.19.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//replace	github.com/qor5/web => ../web
//replace github.com/theplant/docgo => ../../docgo/
