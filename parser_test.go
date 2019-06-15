package parser_test

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	fset := token.NewFileSet() // positions are relative to fset

	src, err := ioutil.ReadFile("testdata/foo.go")
	if err != nil {
		t.Fatal(err)
	}

	// Parse src but stop after processing the imports.
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the imports from the file's AST.
	if strings.HasPrefix(f.Comments[0].Text(), "+build ") {
		fmt.Println(f.Comments[0].Text())
	}
}
