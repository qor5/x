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
					var translationMap = make(map[string]string)
					var isMessage bool
					for _, name := range spec.Names {
						if l, exist := v.LocalesMap[name.Name]; exist {
							locale = l
							for _, messageStructs := range v.RigisterMap[locale] {
								if messageStructs.PkgName == pkgName && messageStructs.StructName == name.Name {
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
							break
						}

						x, ok := unaryExpr.X.(*ast.CompositeLit)
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

							if isMessage {
								translationMap[go_path.Join(pkgName, key.Name)] = strings.Trim(value.Value, "\"")
							}
						}
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
