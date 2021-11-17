package i18n

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/sunfmin/reflectutils"
)

type kv struct {
	key string
	val string
}

type moduleMissing struct {
	missingKeys []kv
	missingVals []kv
}
type DynaBuilder struct {
	lang    string
	missing map[ModuleKey]*moduleMissing
}

func DynaNew() (r *DynaBuilder) {
	return &DynaBuilder{
		missing: make(map[ModuleKey]*moduleMissing),
	}
}

func (d *DynaBuilder) Language(lang string) (r *DynaBuilder) {
	d.lang = lang
	return d
}

func T(req *http.Request, module ModuleKey, key string, args ...string) (r string) {
	return PT(req, module, "", key, args...)
}

func PT(req *http.Request, module ModuleKey, prefix string, key string, args ...string) (r string) {
	defaultVal := strings.NewReplacer(args...).Replace(key)
	msgr := MustGetModuleMessages(req, module, nil)
	if msgr == nil {
		return defaultVal
	}

	var builder *DynaBuilder
	b := req.Context().Value(dynaBuilderKey)
	if b != nil {
		builder = b.(*DynaBuilder)
	}

	fieldKey := strcase.ToCamel(prefix + " " + key)
	val, err := reflectutils.Get(msgr, fieldKey)
	if err != nil {
		if builder != nil {
			builder.putMissingKey(module, fieldKey, key)
		}
		return defaultVal
	}

	if val.(string) == "" {
		if builder != nil {
			builder.putMissingVal(module, fieldKey, key)
		}
	}

	return strings.NewReplacer(args...).Replace(val.(string))
}

func (d *DynaBuilder) putMissingKey(module ModuleKey, key, val string) {
	if d.missing[module] == nil {
		d.missing[module] = &moduleMissing{}
	}
	mm := d.missing[module]

	for _, ck := range mm.missingKeys {
		if ck.key == key {
			return
		}
	}
	mm.missingKeys = append(mm.missingKeys, kv{key, val})
}

func (d *DynaBuilder) putMissingVal(module ModuleKey, key, val string) {
	if d.missing[module] == nil {
		d.missing[module] = &moduleMissing{}
	}
	mm := d.missing[module]

	for _, ck := range mm.missingVals {
		if ck.key == key {
			return
		}
	}
	mm.missingVals = append(mm.missingVals, kv{key, val})
}

func (d *DynaBuilder) HaveMissingKeys() bool {
	return len(d.missing) > 0
}

func (d *DynaBuilder) PrettyMissingKeys() string {
	buf := new(bytes.Buffer)
	for module, missing := range d.missing {

		buf.WriteString(fmt.Sprintf("For module %s, ", module))
		buf.WriteString("Missing the following translations\nCopy these to your Messages struct definition\n============================\n\n")
		for _, kv := range missing.missingKeys {
			_, _ = fmt.Fprintf(buf, "%s string\n", kv.key)
		}
		buf.WriteString("\n")
		buf.WriteString(fmt.Sprintf("Copy these to your Messages struct values for language: `%s`\n\n", d.lang))
		for _, kv := range missing.missingKeys {
			_, _ = fmt.Fprintf(buf, "%s: %#+v,\n", kv.key, kv.val)
		}

		for _, kv := range missing.missingVals {
			_, _ = fmt.Fprintf(buf, "%s: %#+v,\n", kv.key, kv.val)
		}
	}

	return buf.String()
}
