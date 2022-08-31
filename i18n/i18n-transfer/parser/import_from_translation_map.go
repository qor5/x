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

// compare the translation structs and translationsMap
// update the new value from translationsMap
// overwrite new content to files
func ImportFromTranslationsMap(projectPath string, translationsMap map[string]map[string]string) (err error) {
	fset := token.NewFileSet()
	pkgs, err := ParseDir(fset, projectPath, nil, go_parser.AllErrors)
	if err != nil {
		return
	}
	visitor, err := newVisitorAndWalk(fset, pkgs, projectPath)
	if err != nil {
		return
	}

	var isChanged bool
	for path, pkg := range pkgs {
		pkgName := strings.TrimPrefix(path, strings.TrimSuffix(visitor.projectParentPath, "/")+"/")
		for fileName, f := range pkg.Files {
			var isModifiedFile bool
			for _, decl := range f.Decls {
				if decl, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range decl.Specs {
						if spec, ok := spec.(*ast.ValueSpec); ok {
							var isMessage bool
							var locale string
							var structName string
							for _, name := range spec.Names {
								if l, exist := visitor.LocalesMap[name.Name]; exist {
									if _, exist := translationsMap[l]; exist {
										locale = l
										structName = name.Name
										isMessage = true
										break
									}
								}
							}
							if isMessage {
								isMessage = false
								for _, messageStruct := range visitor.RigisterMap[locale] {
									if messageStruct.PkgName == pkgName && messageStruct.StructName == structName {
										isMessage = true
										break
									}
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

								_, ok = x.Type.(*ast.Ident)
								if !ok {
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

`, go_path.Join(fileName, structName, key.Name), value.Value, "\""+translationValue+"\"")
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

			// overwrite new content to file
			if isModifiedFile {
				file, err := os.Create(fileName)
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
