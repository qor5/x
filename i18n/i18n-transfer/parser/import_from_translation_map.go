package parser

import (
	"fmt"
	"go/ast"
	"go/format"
	go_parser "go/parser"
	"go/token"
	"os"
	go_path "path"
	"strings"
)

func ImportFromTranslationsMap(projectPath string, translationsMap map[string]map[string]string) error {
	fset := token.NewFileSet()
	pkgs, err := ParseDir(fset, projectPath, nil, go_parser.AllErrors)
	if err != nil {
		return err
	}
	var isChanged bool
	for path, pkg := range pkgs {
		for fileName, f := range pkg.Files {
			var isModifiedFile bool
			for _, decl := range f.Decls {
				if decl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range decl.Specs {
						if spec, ok := spec.(*ast.ValueSpec); ok {
							var isMessage bool
							var locale string
							for _, name := range spec.Names {
								if _, exist := translationsMap[name.Name]; exist {
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

									if translationValue, exist := translationsMap[locale][getTranslationMapKey(path, key.Name, projectPath)]; isMessage && exist && value.Value != "\""+translationValue+"\"" {
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
										isModifiedFile = true
										isChanged = true
									}
								}
							}
						}
					}
				}
			}

			if isModifiedFile {
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
	if !isChanged {
		fmt.Printf(`
----------------------------------------------
update translation:
	nothing changed
`)
	}
	return nil
}
