package main

type SyntaxKind int

type Token struct {
	kind  SyntaxKind
	value string
}

type NodeOrToken struct {
	node  *Node
	token *Token
}

type Node struct {
	kind     SyntaxKind
	children []*NodeOrToken
	length   int
}

func NewToken(kind SyntaxKind, value string) *Token {
	return &Token{kind: kind, value: value}
}

func NewNode(kind SyntaxKind, children []*NodeOrToken) *Node {

	length := 0

	for _, child := range children {
		length += child.Len()
	}

	return &Node{kind: kind, children: []*NodeOrToken{}, length: length}
}

func (t *Token) Len() int {
	return len(t.value)
}

func (t *Token) Kind() SyntaxKind {
	return t.kind
}

func (t *Token) Value() string {
	return t.value
}

func (n *Node) Len() int {
	return n.length
}

func (n *Node) Kind() SyntaxKind {
	return n.kind
}

func (nt *NodeOrToken) Len() int {
	if nt.node != nil {
		return nt.node.Len()
	}
	return nt.token.Len()
}

func (nt *NodeOrToken) Kind() SyntaxKind {
	if nt.node != nil {
		return nt.node.Kind()
	}
	return nt.token.Kind()
}

func (n *Node) Children() []*NodeOrToken {
	return n.children
}

func (nt *NodeOrToken) IsToken() bool {
	return nt.node == nil
}

func (nt *NodeOrToken) IsNode() bool {
	return nt.node != nil
}

const (
	EOF SyntaxKind = iota
)
