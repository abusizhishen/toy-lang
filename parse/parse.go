package parse

import (
	"github.com/abusizhishen/toy-lang/ast"
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
	var p = &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) NextStatement() ast.Statement {
	tok := p.l.NextToken()

	switch tok {

	}

	return ast.Statement{}
}

func (p *Parser) Parse() {
	for {
		tok := p.l.NextToken()
	}
}
