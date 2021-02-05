package lexer

import "github.com/abusizhishen/toy-lang/token"

type Lexer struct {
	position     int
	readPosition int
	length       int
	Str          string
	ch           byte
}

func NewLexer(str string) *Lexer {
	return &Lexer{Str: str, length: len(str)}
}

func (l *Lexer) readCh() {
	if l.readPosition >= l.length {
		l.ch = 0
	} else {
		l.ch = l.Str[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readNextCh() byte {
	if l.readPosition >= l.length {
		return 0
	} else {
		return l.Str[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	l.skipSpace()
	l.readCh()
	var tok token.Token
	switch l.ch {
	case '+':
		tok = token.Token{Type: token.Add, Literal: string(l.ch)}
	case '-':
		tok = token.Token{Type: token.Sub, Literal: string(l.ch)}
	case '*':
		tok = token.Token{Type: token.Mul, Literal: string(l.ch)}
	case '/':
		tok = token.Token{Type: token.Div, Literal: string(l.ch)}
	case '(':
		tok = token.Token{Type: token.LYkh, Literal: string(l.ch)}
	case ')':
		tok = token.Token{Type: token.RYkh, Literal: string(l.ch)}
	case '{':
		tok = token.Token{Type: token.LHkh, Literal: string(l.ch)}
	case '}':
		tok = token.Token{Type: token.RHkh, Literal: string(l.ch)}
	case '[':
		tok = token.Token{Type: token.LFkh, Literal: string(l.ch)}
	case ']':
		tok = token.Token{Type: token.RFkh, Literal: string(l.ch)}
	case '!':
		if l.readNextCh() == '=' {
			l.readCh()
			tok = token.Token{Type: token.NEQ, Literal: "!="}
			break
		}
		tok = token.Token{Type: token.Not, Literal: string(l.ch)}
	case '=':
		if l.readNextCh() == '=' {
			l.readCh()
			tok = token.Token{Type: token.EQ, Literal: "=="}
			break
		}
		tok = token.Token{Type: token.Assign, Literal: string(l.ch)}
	case '>':
	case ';':
		tok = token.Token{Type: token.FH, Literal: string(l.ch)}
	case '"':
		tok = l.readString()
	case '|':
		if l.readNextCh() == '|' {
			l.readCh()
			tok = token.Token{Type: token.Or, Literal: "||"}
			break
		}
		tok = token.Token{Type: token.EOF, Literal: string(l.ch)}
	case '&':
		if l.readNextCh() == '&' {
			l.readCh()
			tok = token.Token{Type: token.And, Literal: "&&"}
			break
		}
		tok = token.Token{Type: token.EOF, Literal: string(l.ch)}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: string(l.ch)}

	default:
		if isWord(l.ch) {
			tok = l.readId()
			break
		}

		if isNum(l.ch) {
			tok = l.readNum()
			break
		}

		tok = token.Token{Type: token.InvalidToken, Literal: string(l.ch)}
	}

	return tok
}

func (l *Lexer) readId() token.Token {
	var position = l.position
	for {
		if !isWord(l.ch) {
			break
		}
		l.readCh()
	}

	var literal = l.Str[position:l.position]
	if _, ok := token.KeyWordMap[literal]; ok {
		return token.Token{Type: token.KeyWord, Literal: literal}
	}
	return token.Token{Type: token.Identify, Literal: literal}
}

func (l *Lexer) readNum() token.Token {
	var position = l.position
	for {
		if !isNum(l.ch) {
			break
		}
		l.readCh()
	}

	return token.Token{Type: token.Num, Literal: l.Str[position:l.position]}
}

func (l *Lexer) readString() token.Token {
	var position = l.position
	l.readCh()

	for {
		switch l.ch {
		case '"':
			l.position++
			goto end
		case 0:
			return token.Token{Type: token.InvalidToken, Literal: l.Str[position:l.position]}
		case '\\':
			if l.readNextCh() == '"' {
				l.readCh()
			} else {
				continue
			}
		}
		l.readCh()
	}

end:

	return token.Token{Type: token.Str, Literal: l.Str[position:l.position]}
}

func (l *Lexer) skipSpace() {
	for {
		switch l.readNextCh() {
		case ' ', '\n', '\r', '\t':
			l.readCh()
		default:
			return
		}
	}
}

func isWord(ch byte) bool {
	return ch >= 'A' && ch <= 'z' || '_' == ch
}

func isNum(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
