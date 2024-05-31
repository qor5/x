package parser

import (
	"go/ast"
	"go/token"
	"strings"
)

type Visitor struct {
	// map[locale][]MessageStruct
	// example:
	// {
	//      "English": [
	//              {
	//                      "PkgName": "x/i18n",
	//                      "StructName": "Messages_en_US"
	//              }
	//      ],
	// }
	RigisterMap map[string][]MessageStruct

	// map[locale][]MessageStruct
	// example:
	// {
	//      "Messages_en_US": "English",
	//      "Messages_zh_CN": "SimplifiedChinese"
	// }
	LocalesMap map[string]string

	// map[pkgName][]structs
	// example:
	// {
	//      "x/i18n": []*ast.GenDecl,
	// }
	Variables map[string][]*ast.GenDecl

	// the current package path
	// example:
	// when visit file "x/i18n/i18n_test.go", currentPkgPath = "x/i18n"
	currentPkgPath string

	// the current import map
	// example:
	// when visit file "x/i18n/i18n_test.go", currentImportMap =
	// {
	//      "fmt": "fmt",
	//      "http": "net/http",
	//      "httptest": "net/http/httptest",
	//      "strings": "strings",
	//      "testing": "testing"
	//      "i18n": "github.com/qor5/x/v3/i18n",
	//      "testingutils": "github.com/theplant/testingutils",
	//      "language": "golang.org/x/text/language",
	// }
	currentImportMap  map[string]string
	projectParentPath string
	fset              *token.FileSet
}

type MessageStruct struct {
	PkgName    string
	StructName string
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	// find all global variables and insert it into v.Variables
	if f, ok := node.(*ast.File); ok {
		temp := strings.Split(v.fset.File(f.Package).Name(), "/")
		pkgName := strings.TrimPrefix(strings.Join(temp[:len(temp)-1], "/"), strings.TrimSuffix(v.projectParentPath, "/")+"/")
		for _, decl := range f.Decls {
			if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
				v.Variables[pkgName] = append(v.Variables[pkgName], genDecl)
			}
		}
	}

	// find all places calling func RegisterForModule
	// analyze the calls and fill it into v.LocalesMap and v.RigisterMap
	if callExpr, ok := node.(*ast.CallExpr); ok {
		if selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok && selectorExpr.Sel.Name == "RegisterForModule" {
			if len(callExpr.Args) == 3 {
				if selectorExpr2, ok := callExpr.Args[0].(*ast.SelectorExpr); ok {
					if ident, ok := selectorExpr2.X.(*ast.Ident); ok && ident.Name == "language" {
						if ident2, ok := callExpr.Args[2].(*ast.Ident); ok {
							var messageStruct MessageStruct
							messageStruct.PkgName = strings.TrimPrefix(v.currentPkgPath, strings.TrimSuffix(v.projectParentPath, "/")+"/")
							messageStruct.StructName = ident2.Name

							v.LocalesMap[messageStruct.StructName] = selectorExpr2.Sel.Name
							v.RigisterMap[selectorExpr2.Sel.Name] = append(v.RigisterMap[selectorExpr2.Sel.Name], messageStruct)
						}
						if selectorExpr3, ok := callExpr.Args[2].(*ast.SelectorExpr); ok {
							var messageStruct MessageStruct
							messageStruct.PkgName = v.currentImportMap[selectorExpr3.X.(*ast.Ident).Name]
							messageStruct.StructName = selectorExpr3.Sel.Name

							v.LocalesMap[messageStruct.StructName] = selectorExpr2.Sel.Name
							v.RigisterMap[selectorExpr2.Sel.Name] = append(v.RigisterMap[selectorExpr2.Sel.Name], messageStruct)
						}
					}
				}
			}
		}
	}

	return v
}

// declare a Visitor
// walk all files and fill the RigisterMap, LocalesMap and Variables
func newVisitorAndWalk(fset *token.FileSet, pkgs map[string]*ast.Package, projectPath string) (v *Visitor, err error) {
	v = &Visitor{
		RigisterMap:      make(map[string][]MessageStruct),
		LocalesMap:       make(map[string]string),
		Variables:        make(map[string][]*ast.GenDecl),
		currentImportMap: make(map[string]string),
		fset:             fset,
	}
	for pkgPath, pkg := range pkgs {
		for _, f := range pkg.Files {
			for _, decl := range f.Decls {
				if decl, ok := decl.(*ast.GenDecl); ok {
					if decl.Tok == token.IMPORT {
						for _, spec := range decl.Specs {
							if spec, ok := spec.(*ast.ImportSpec); ok {
								var importName string
								var importValue string

								if spec.Name == nil {
									temp := strings.Split(strings.Trim(spec.Path.Value, "\""), "/")
									importName = temp[len(temp)-1]
								} else {
									importName = spec.Name.Name
								}

								importValue = strings.Trim(spec.Path.Value, "\"")
								v.currentImportMap[importName] = importValue
							}
						}
					}
				}
			}
			v.currentPkgPath = pkgPath
			temp := strings.Split(projectPath, "/")
			v.projectParentPath = strings.Join(temp[:len(temp)-1], "/")

			ast.Walk(v, f)
		}
	}
	return
}
