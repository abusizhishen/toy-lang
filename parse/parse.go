package parse

import (
	"github.com/abusizhishen/toy-lang/lexer"
	"github.com/abusizhishen/toy-lang/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	Err          []string
}

func New(l *lexer.Lexer) *Parser {
	var p = &Parser{
		l:            l,
		currentToken: token.Token{},
		peekToken:    token.Token{},
		Err:          nil,
	}

	return p
}

func (p *Parser) NextToken() {

}

func (p *Parser) Parse() {

}
