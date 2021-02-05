package ast

type Statements []*Statement

type Statement interface {
	Node
	Literal() string
}

type AssignStatement struct {
}

func (a AssignStatement) Literal() string {
	return ""
}

type ReturnStatement struct {
}

func (r ReturnStatement) Literal() string {
	return ""
}

type Node interface {
}

func New() {

}
