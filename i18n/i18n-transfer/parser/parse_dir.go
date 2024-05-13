package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	go_path "path"
	"path/filepath"
	"strings"
)

func ParseDir(fset *token.FileSet, path string, filter func(fs.FileInfo) bool, mode parser.Mode) (pkgs map[string]*ast.Package, first error) {
	list, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	pkgs = make(map[string]*ast.Package)
	for _, d := range list {
		if d.IsDir() {
			insidePkgs := make(map[string]*ast.Package)
			insidePkgs, first = ParseDir(fset, go_path.Join(path, d.Name()), filter, mode)
			if first != nil {
				return
			}
			for s, a2 := range insidePkgs {
				pkgs[s] = a2
			}
		}
		if !strings.HasSuffix(d.Name(), ".go") {
			continue
		}
		if filter != nil {
			info, err := d.Info()
			if err != nil {
				return nil, err
			}
			if !filter(info) {
				continue
			}
		}
		filename := filepath.Join(path, d.Name())
		if src, err := parser.ParseFile(fset, filename, nil, mode); err == nil {
			name := src.Name.Name
			pkg, found := pkgs[path]
			if !found {
				pkg = &ast.Package{
					Name:  name,
					Files: make(map[string]*ast.File),
				}
				pkgs[path] = pkg
			}
			pkg.Files[filename] = src
		} else if first == nil {
			first = err
		}
	}

	return
}
