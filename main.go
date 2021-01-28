package main

import (
	"fmt"
	"github.com/abusizhishen/toy-lang/parse"
)

func main() {
	var s = "abcdefg{("
	var lex = parse.NewLexer(s)

	for {
		token := lex.NextToken()

		if token.Type == parse.EOF{
			break
		}

		fmt.Printf("token:%+v\n", token)
	}
}
