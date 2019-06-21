package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"

//	"github.com/davecgh/go-spew/spew"
)

func main() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3; var s = \"sss\"", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

//	fmt.Println(f)
//	spew.Dump(f)

	printer.Fprint(os.Stdout, fs, f)

	var v visitor
	ast.Walk(v, f)

}

type visitor struct{}

func (v visitor) Visit(n ast.Node) ast.Visitor {
//	fmt.Printf("%T\n", n)
	switch d := n.(type) {
	case *ast.BasicLit:
		if(d.Kind == token.STRING) {
			fmt.Printf("string: ")
		}
		fmt.Printf("%s\n",d.Value)
	}
	return v
}
