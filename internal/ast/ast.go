package ast

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode() // marker method
}

type Expression interface {
	Node
	expressionNode() // marker method
}
