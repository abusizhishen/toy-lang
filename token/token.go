package token

type Token struct {
	Type    Type
	Literal string
}

type Type string

const (
	InvalidToken Type = "invalid token"
	Id           Type = "id"
	Num          Type = "number"
	EOF          Type = "EOF"
	Str          Type = "string"
	EQ           Type = "Type = Type = "
	Assign       Type = "Type = "
	LHkh         Type = "{"
	RHkh         Type = "}"
	LYkh         Type = "("
	RYkh         Type = ")"
	LFkh         Type = "("
	RFkh         Type = ")"
	FH           Type = ";"
	Not          Type = "!"
	NEQ          Type = "!Type = "
	And          Type = "&&"
	Or           Type = "||"
	Refer        Type = "&"
	Add          Type = "+"
	Sub          Type = "-"
	Mul          Type = "*"
	Div          Type = "/"
	KeyWord           = "keyword"
)

var KeyWordMap = map[string]string{
	"var":    "",
	"int":    "",
	"let":    "",
	"fun":    "",
	"type":   "",
	"switch": "",
	"return": "",
}
