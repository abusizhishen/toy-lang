package main

import (
	"fmt"
	"github.com/abusizhishen/toy-lang/parse"
)

func main() {
	var s = ""
	var lex = parse.NewLexer(s)

	for {
		token,err := lex.NextToken()
		if err != nil{
			fmt.Printf("parse err:%s",err)
			return
		}

		fmt.Printf("token:%+v", token)
	}
}
