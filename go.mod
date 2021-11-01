module github.com/goplaid/x

go 1.16

require (
	github.com/Pallinder/go-randomdata v1.2.0
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/go-chi/chi v1.5.4
	github.com/google/uuid v1.3.0 // indirect
	github.com/goplaid/multipartestutils v0.0.3
	github.com/goplaid/web v1.1.18-0.20211101151911-49372dcb059c
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/iancoleman/strcase v0.2.0
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/inflection v1.0.0
	github.com/juju/ansiterm v0.0.0-20210706145210-9283cdf370b5 // indirect
	github.com/lib/pq v1.10.3 // indirect
	github.com/lunixbochs/vtclean v1.0.0 // indirect
	github.com/manifoldco/promptui v0.8.0
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/microcosm-cc/bluemonday v1.0.15 // indirect
	github.com/ory/ladon v1.2.0
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/xid v1.3.0
	github.com/russross/blackfriday v1.6.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/shurcooL/github_flavored_markdown v0.0.0-20210228213109-c3a9aa474629
	github.com/shurcooL/go v0.0.0-20200502201357-93f07166e636 // indirect
	github.com/shurcooL/go-goon v0.0.0-20210110234559-7585751d9a17 // indirect
	github.com/shurcooL/highlight_diff v0.0.0-20181222201841-111da2e7d480 // indirect
	github.com/shurcooL/highlight_go v0.0.0-20191220051317-782971ddf21b // indirect
	github.com/shurcooL/octicon v0.0.0-20191102190552-cbb32d6a785c // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d // indirect
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e // indirect
	github.com/sunfmin/reflectutils v1.0.2
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
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf // indirect
	golang.org/x/sys v0.0.0-20210921065528-437939a70204 // indirect
	golang.org/x/text v0.3.7
	gorm.io/driver/postgres v1.1.1
	gorm.io/driver/sqlite v1.1.5
	gorm.io/gorm v1.21.15
)

//replace	github.com/goplaid/web => ../web
