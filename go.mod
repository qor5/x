module github.com/goplaid/x

go 1.17

require (
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/google/go-cmp v0.5.9
	github.com/google/uuid v1.3.0
	github.com/goplaid/web v1.2.2
	github.com/iancoleman/strcase v0.2.0
	github.com/jinzhu/inflection v1.0.0
	github.com/lib/pq v1.10.3
	github.com/manifoldco/promptui v0.9.0
	github.com/markbates/goth v1.75.1
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/ory/ladon v1.2.0
	github.com/pquerna/otp v1.3.0
	github.com/spf13/cast v1.5.0
	github.com/stretchr/testify v1.8.0
	github.com/sunfmin/reflectutils v1.0.3
	github.com/theplant/htmlgo v1.0.3
	github.com/theplant/testingutils v0.0.0-20190603093022-26d8b4d95c61
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	golang.org/x/text v0.3.7
	gorm.io/driver/postgres v1.4.5
	gorm.io/driver/sqlite v1.4.3
	gorm.io/gorm v1.24.1-0.20221019064659-5dd2bb482755
)

require (
	cloud.google.com/go v0.67.0 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/boombuler/barcode v1.0.1-0.20190219062509-6c824513bacc // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/go-playground/form v3.1.4+incompatible // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.6.2 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.1.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/mrjones/oauth v0.0.0-20180629183705-f4e24b6d100c // indirect
	github.com/ory/pagination v0.0.1 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//replace	github.com/goplaid/web => ../web
//replace github.com/theplant/docgo => ../../docgo/
