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

	var statement ast.Statement
	switch tok.Type {
	case token.RETURN:
		return p.readAssignStatement()
	case token.Identify:
		if p.peekToken.Type == token.EQ {
			statement = p.readAssignStatement()
			break
		}

	}

	return statement
}

func (p *Parser) Parse() *ast.Statements {
	var stmts = new(ast.Statements)
	for {
		if p.currentToken.Type == token.EOF {
			break
		}
	}

	return stmts
}

func (p *Parser) readAssignStatement() ast.Statement {

	return ast.AssignStatement{}
}

func (p *Parser) readReturnStatement() ast.Statement {

	return ast.ReturnStatement{}
}
