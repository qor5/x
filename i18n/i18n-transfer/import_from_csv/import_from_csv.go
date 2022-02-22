package import_from_csv

import (
	"fmt"
	"go/ast"
	"go/format"
	go_parser "go/parser"
	"go/token"
	"os"
	go_path "path"
	"strings"

	"github.com/goplaid/x/i18n/i18n-transfer/parser"
)

func ImportFromCsv(dir string, translationMap map[string]map[string]string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, go_parser.AllErrors)
	if err != nil {
		return err
	}
	for path, pkg := range pkgs {
		for fileName, f := range pkg.Files {
			var isModiyiedFile bool
			for _, decl := range f.Decls {
				if decl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range decl.Specs {
						if spec, ok := spec.(*ast.ValueSpec); ok {
							var isMessage bool
							var locale string
							for _, name := range spec.Names {
								if _, exist := translationMap[name.Name]; exist {
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

									if translationValue, exist := translationMap[locale][go_path.Join(path, key.Name)]; isMessage && exist && value.Value != "\""+translationValue+"\"" {
										fmt.Printf(`
----------------------------------------------
update translation:
	%s
from:
	%s
to:
	%s
----------------------------------------------

`, go_path.Join(fileName, locale, key.Name), value.Value, "\""+translationValue+"\"")
										value.Value = "\"" + translationValue + "\""
										isModiyiedFile = true
									}
								}
							}
						}
					}
				}
			}

			if isModiyiedFile {
				file, err := os.OpenFile(fileName, os.O_WRONLY, 0)
				defer file.Close()
				if err != nil {
					return err
				}
				err = format.Node(file, fset, f)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
