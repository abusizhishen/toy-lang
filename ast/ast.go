package ast

type Ast struct {
	ss []*Statement
}

type Statement interface {
	Literal() string
}
