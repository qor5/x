module github.com/goplaid/x

go 1.16

require (
	github.com/Pallinder/go-randomdata v1.2.0
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/go-chi/chi v1.5.4
	github.com/google/uuid v1.3.0 // indirect
	github.com/goplaid/multipartestutils v0.0.3
	github.com/goplaid/web v1.1.23
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/iancoleman/strcase v0.2.0
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/inflection v1.0.0
	github.com/lib/pq v1.10.3 // indirect
	github.com/manifoldco/promptui v0.9.0
	github.com/ory/ladon v1.2.0
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/xid v1.3.0
	github.com/russross/blackfriday v1.6.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0
	github.com/sunfmin/reflectutils v1.0.3
	github.com/theplant/docgo v0.0.7
	github.com/theplant/gofixtures v1.1.0
	github.com/theplant/htmlgo v1.0.3
	github.com/theplant/testingutils v0.0.0-20190603093022-26d8b4d95c61
	github.com/thoas/go-funk v0.9.1
	github.com/yosssi/gohtml v0.0.0-20201013000340-ee4748c638f4
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1
	goji.io v2.0.2+incompatible
	golang.org/x/crypto v0.0.0-20210920023735-84f357641f63 // indirect
	golang.org/x/sys v0.0.0-20210921065528-437939a70204 // indirect
	golang.org/x/text v0.3.7
	gorm.io/driver/postgres v1.1.1
	gorm.io/driver/sqlite v1.1.5
	gorm.io/gorm v1.21.15
)

//replace	github.com/goplaid/web => ../web
//replace github.com/theplant/docgo => ../../docgo/
