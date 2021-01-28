package parse

type Token struct {
	Type TokenType
	Literal string
}

type TokenType string

