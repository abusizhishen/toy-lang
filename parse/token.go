package parse

type Token struct {
	Type TokenType
	Literal string
}

type TokenType int
const (
	InvalidToken = iota + 1
	Id
	Num
	EOF
)
