package export_to_csv

import (
	"go/ast"
	go_parser "go/parser"
	"go/token"
	go_path "path"
	"strings"

	"github.com/goplaid/x/i18n/i18n-transfer/parser"
)

func GetTranslationsMap(projectDir string) (translations map[string]map[string]string, err error) {
	translations = make(map[string]map[string]string)
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, projectDir, nil, go_parser.AllErrors)
	if err != nil {
		return
	}

	for path, pkg := range pkgs {
		for _, f := range pkg.Files {
			for _, decl := range f.Decls {
				if decl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range decl.Specs {
						if spec, ok := spec.(*ast.ValueSpec); ok {
							var locale string
							var translation = make(map[string]string)
							var isMessage bool
							for _, name := range spec.Names {
								if strings.Contains(name.Name, "Message") {
									locale = name.Name
									isMessage = true
								}
							}
							if !isMessage {
								continue
							}
							for _, values := range spec.Values {
								unaryExpr, ok := values.(*ast.UnaryExpr)
								if !ok {
									isMessage = false
									break
								}

								x, ok := unaryExpr.X.(*ast.CompositeLit)
								if !ok {
									isMessage = false
									break
								}

								xType, ok := x.Type.(*ast.Ident)
								if !ok || !strings.Contains(xType.Name, "Message") {
									isMessage = false
									break
								}

								for _, elt := range x.Elts {
									keyValueExpr, ok := elt.(*ast.KeyValueExpr)
									if !ok {
										isMessage = false
										break
									}

									key, ok := keyValueExpr.Key.(*ast.Ident)
									if !ok {
										isMessage = false
										break
									}

									value, ok := keyValueExpr.Value.(*ast.BasicLit)
									if !ok {
										isMessage = false
										break
									}

									if isMessage {
										translation[go_path.Join(path, key.Name)] = strings.Trim(value.Value, "\"")
									}
								}
							}

							if isMessage {
								if translations[locale] == nil {
									translations[locale] = make(map[string]string)
								}
								for k, v := range translation {
									translations[locale][k] = v
								}
							}
						}
					}
				}
			}
		}
	}
	return
}
