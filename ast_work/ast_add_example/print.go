package ast_add_example

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
)

func PrintAst() {
	fset := token.NewFileSet()
	fpath, _ := filepath.Abs("./templ.go")
	astF, err := parser.ParseFile(fset, fpath, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("解析文件错误:", err)
		return
	}

	//fmt.Println(astF)
	ast.Print(fset, astF)
}
