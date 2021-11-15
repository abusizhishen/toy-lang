package main

import (
	"fmt"
	"github.com/abusizhishen/toy-lang/lexer"
	"github.com/abusizhishen/toy-lang/token"
)

func main() {
	var s = `var m = 1;
	var b =2;
a+b
`
	var lex = lexer.NewLexer(s)

	for {
		tok := lex.NextToken()

		if tok.Type == token.EOF {
			break
		}

		fmt.Printf("token:%+v\n", tok)
	}
}
