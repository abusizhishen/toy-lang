package parse

type Token struct {
	Type TokenType
	Literal string
}

type TokenType string
const (
	InvalidToken = "invalid token"
	Id = "id"
	Num = "number"
	EOF = "EOF"
	Str = "string"
	EQ = "=="
	Assign = "="
	LHkh = "{"
	RHkh = "}"
	LYkh = "("
	RYkh = ")"
	FH = ";"

)

var keyWordMap = map[string]string{
	"int":"",
	"let":"",
	"fun":"",
	"type":"",
	"switch":"",
	"return":"",
}