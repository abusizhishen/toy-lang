package parse

type Lexer struct {
	position int
	nextPosition int
	length int
	Str string
}

func NewLexer(str string) *Lexer {
	return &Lexer{Str: str}
}

func (l *Lexer)readNextCh() byte {
	if l.nextPosition >= l.length{
		return '0'
	}
	return l.Str[l.nextPosition]
}

func (l *Lexer) NextToken() (Token,error){
	switch l.readNextCh() {
	case '{':

	}
	return Token{},nil
}

