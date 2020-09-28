package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

func main() {

	src, err := ioutil.ReadFile("./temp.go")
	if err != nil {
		panic(err)
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	typeDecl := f.Decls[0].(*ast.GenDecl)
	structDecl := typeDecl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
	fields := structDecl.Fields.List

	for _, field := range fields {

		tag := field.Tag.Value
		tag2 := field.Tag.ValuePos
		fmt.Println(tag)
		fmt.Println(tag2)

		typeNameExpr := field.Names
		fmt.Println(typeNameExpr[0].String())

		typeExpr := field.Type

		start := typeExpr.Pos() - 1
		end := typeExpr.End() - 1

		// grab it in source
		typeInSource := string(src[start:end])
		fmt.Println(typeInSource)
	}

}
