package parser

import (
	"go/ast"
	go_parser "go/parser"
	"go/token"
	go_path "path"
	"strings"
)

// get translationsMap from projectPath
func ExportToTranslationsMap(projectPath string) (translationsMap map[string]map[string]string, err error) {
	fset := token.NewFileSet()
	pkgs, err := ParseDir(fset, projectPath, nil, go_parser.AllErrors)
	if err != nil {
		return
	}
	v, err := newVisitorAndWalk(fset, pkgs, projectPath)
	if err != nil {
		return
	}

	return getTranslationsMapFromVistor(v), nil
}

// traverse all global variables and check if it is in v.RigisterMap
// if true, transfer it into translationsMap
func getTranslationsMapFromVistor(v *Visitor) map[string]map[string]string {
	translationsMap := make(map[string]map[string]string)

	for pkgName, structs := range v.Variables {
		for _, astruct := range structs {
			for _, spec := range astruct.Specs {
				if spec, ok := spec.(*ast.ValueSpec); ok {
					var locale string
					translationMap := make(map[string]string)
					var isMessage bool
					for _, name := range spec.Names {
						if l, exist := v.LocalesMap[name.Name]; exist {
							locale = l
							for _, messageStructs := range v.RigisterMap[locale] {
								if strings.HasSuffix(messageStructs.PkgName, pkgName) && messageStructs.StructName == name.Name {
									isMessage = true
									break
								}
							}
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
							continue
						}
						isMessage = v.translationExport(translationMap, pkgName, unaryExpr.X)
					}

					if isMessage {
						if translationsMap[locale] == nil {
							translationsMap[locale] = make(map[string]string)
						}
						for k, v := range translationMap {
							translationsMap[locale][k] = v
						}
					}
				}
			}
		}
	}
	return translationsMap
}

func (v *Visitor) translationExport(translationMap map[string]string, pkgName string, x interface{}) bool {
	_, ok := x.(*ast.CompositeLit)
	if !ok {
		return false
	}
	for _, elt := range x.(*ast.CompositeLit).Elts {
		keyValueExpr, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			return false
		}

		key, ok := keyValueExpr.Key.(*ast.Ident)
		if !ok {
			return false
		}

		value, ok := keyValueExpr.Value.(*ast.BasicLit)
		if ok {
			translationMap[go_path.Join(pkgName, key.Name)] = strings.Trim(value.Value, "\"")
		} else {
			// embed struct
			if embed, ok := keyValueExpr.Value.(*ast.Ident); ok {
				// struct is in other file
				if embed.Obj == nil {
					for name, dels := range v.Variables {
						if pkgName != name {
							continue
						}
						for _, del := range dels {
							for _, spec := range del.Specs {
								if valueSpec, ok := spec.(*ast.ValueSpec); ok {
									for _, specName := range valueSpec.Names {
										if specName.Name != embed.Name {
											continue
										}
										embed = specName
										goto JUMP
									}
								}
							}
						}
					}
				}
			JUMP:
				if del, ok := embed.Obj.Decl.(*ast.ValueSpec); ok {
					for _, val := range del.Values {
						v.translationExport(translationMap, pkgName, val)
					}
				}
			}
		}
	}
	return true
}
