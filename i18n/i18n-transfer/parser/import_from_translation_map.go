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
							if !isMessage {
								continue
							}

							isMessage = false
							for _, messageStruct := range visitor.RigisterMap[locale] {
								if strings.HasSuffix(messageStruct.PkgName, pkgName) && messageStruct.StructName == structName {
									isMessage = true
									break
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

								modified, changed := visitor.translationImport(pkgs, fset, projectPath, translationsMap, unaryExpr.X, pkgName, locale, path, fileName, structName)
								if modified {
									isModifiedFile = modified
								}
								if changed {
									isChanged = changed
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

func (v *Visitor) translationImport(pkgs map[string]*ast.Package, fset *token.FileSet, projectPath string, translationsMap map[string]map[string]string, x interface{}, pkgName string, locale string, path string, fileName string, structName string) (bool, bool) {
	var isModifiedFile, isChanged bool = false, false

	_, ok := x.(*ast.CompositeLit)
	if !ok {
		return false, false
	}

	for _, elt := range x.(*ast.CompositeLit).Elts {
		keyValueExpr, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			return false, false
		}

		key, ok := keyValueExpr.Key.(*ast.Ident)
		if !ok {
			return false, false
		}

		value, ok := keyValueExpr.Value.(*ast.BasicLit)
		if !ok {
			var needWriteFile bool = false
			var writeF *ast.File
			var writeFileName string
			if embed, ok := keyValueExpr.Value.(*ast.Ident); ok {
				if embed.Obj == nil {
					// other file
					for pkgPath, pkg := range pkgs {
						// same pkg
						if pkgName != strings.TrimPrefix(pkgPath, strings.TrimSuffix(v.projectParentPath, "/")+"/") {
							continue
						}
						for fName, f := range pkg.Files {
							// skip same file
							if fName == fileName {
								continue
							}
							for _, decl := range f.Decls {
								if decl, ok := decl.(*ast.GenDecl); ok {
									for _, spec := range decl.Specs {
										if spec, ok := spec.(*ast.ValueSpec); ok {
											for _, specName := range spec.Names {
												if specName.Name == embed.Name {
													embed = specName
													needWriteFile = true
													writeF = f
													writeFileName = fName
													goto JUMP
												}
											}
										}
									}
								}
							}
						}
					}
				}
			JUMP:
				if del, ok := embed.Obj.Decl.(*ast.ValueSpec); ok {
					for _, val := range del.Values {
						modified, changed := v.translationImport(pkgs, fset, projectPath, translationsMap, val, pkgName, locale, path, fileName, structName)
						if modified {
							isModifiedFile = modified
						}
						if changed {
							isChanged = changed
						}
					}
				}
				if needWriteFile {
					// if changed other file, need write it
					file, err := os.Create(writeFileName)
					defer file.Close()
					if err != nil {
						panic(err)
					}

					err = format.Node(file, fset, writeF)
					if err != nil {
						panic(err)
					}
				}
			}
			continue
		}
		modified, changed := v.importValue(projectPath, translationsMap, locale, path, key, value, fileName, structName)
		if modified {
			isModifiedFile = modified
		}
		if changed {
			isChanged = changed
		}
	}
	return isModifiedFile, isChanged
}

func (v *Visitor) importValue(projectPath string, translationsMap map[string]map[string]string, locale string, path string, key *ast.Ident, value *ast.BasicLit, fileName string, structName string) (bool, bool) {
	if translationValue, exist := translationsMap[locale][getTranslationMapKey(path, key.Name, projectPath)]; exist && value.Value != "\""+translationValue+"\"" {
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
		return true, true
	}
	return false, false
}
