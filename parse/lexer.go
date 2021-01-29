package parse

import "fmt"

type Lexer struct {
	position int
	readPosition int
	length int
	Str string
	ch byte
}

func NewLexer(str string) *Lexer {
	return &Lexer{Str: str, length: len(str)}
}

func (l *Lexer)readCh()  {
	if l.readPosition >= l.length{
		l.ch = 0
	}else {
		l.ch = l.Str[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition +=1
}

func (l *Lexer) NextToken() (Token){
	l.readCh()
	switch l.ch {
	case '{':
	case '=':
	case '>':
	case ';':
	case 0:
		return Token{
			Type:    EOF,
			Literal: string(l.ch),
		}
		
	default:
		if isWord(l.ch) {
			return l.readKey()
		}

		if isNum(l.ch) {
			return l.readNum()
		}

		return Token{
			Type:    InvalidToken,
			Literal: string(l.ch),
		}
	}
	return Token{}
}

func (l *Lexer)readKey() Token {
	var position = l.position
	for {
		l.readCh()

		if !isWord(l.ch){
			break
		}
	}

	return Token{
		Type:    Id,
		Literal: l.Str[position:l.readPosition],
	}
}

func (l *Lexer)readNum() Token {
	var position = l.position
	for {
		l.readCh()

		fmt.Printf("ch:%v, idx:%d\n", l.ch,l.position)
		if !isNum(l.ch){
			break
		}
	}

	return Token{
		Type:    Num,
		Literal: l.Str[position:l.readPosition],
	}
}

func isWord(ch byte) bool {
	return ch >= 'a' && ch <= 'Z' || '_' == ch
}

func isNum(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
