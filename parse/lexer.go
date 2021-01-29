package parse

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

func (l *Lexer)readNextCh() byte {
	if l.readPosition >= l.length{
		return 0
	}else {
		return l.Str[l.readPosition]
	}
}

func (l *Lexer) NextToken() Token{
	l.skipSpace()
	l.readCh()
	var tok Token
	switch l.ch {
	case '{':
		tok = Token{
			Type:    LHkh,
			Literal: string(l.ch),
		}
	case '=':
		tok = Token{
			Type:    Assign,
			Literal: string(l.ch),
		}
	case '>':
	case ';':
		tok = Token{
			Type:    FH,
			Literal: string(l.ch),
		}
	case '"':
		tok = l.readString()

	case 0:
		tok = Token{
			Type:    EOF,
			Literal: string(l.ch),
		}
		
	default:
		if isWord(l.ch) {
			tok = l.readId()
			break
		}

		if isNum(l.ch) {
			tok = l.readNum()
			break
		}

		tok =  Token{
			Type:    InvalidToken,
			Literal: string(l.ch),
		}
	}

	return tok
}

func (l *Lexer)readId() Token {
	var position = l.position
	for {
		if !isWord(l.ch){
			break
		}
		l.readCh()
	}

	var literal = l.Str[position:l.position]
	if _,ok := keyWordMap[literal];ok{
		return Token{
			Type:    TokenType(literal),
			Literal: literal,
		}
	}
	return Token{
		Type:    Id,
		Literal: literal,
	}
}

func (l *Lexer)readNum() Token {
	var position = l.position
	for {
		if !isNum(l.ch){
			break
		}
		l.readCh()
	}

	return Token{
		Type:    Num,
		Literal: l.Str[position:l.position],
	}
}

func (l *Lexer)readString() Token {
	var position = l.position
	l.readCh()

	for {
		switch l.ch {
		case '"':
			l.position++
			goto end
		case 0:
			return Token{
				Type:    InvalidToken,
				Literal: l.Str[position:l.position],
			}
		case '\\':
			if l.readNextCh() == '"'{
				l.readCh()
			}else {
				continue
			}
		}
		l.readCh()
	}

	end:

	return Token{
		Type:   Str ,
		Literal: l.Str[position:l.position],
	}
}

func (l *Lexer) skipSpace() {
	for {
		switch l.readNextCh() {
		case ' ', '\n','\r','\t':
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
