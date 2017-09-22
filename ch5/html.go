package html

type Node struct {
	Type						NodeType
	Data						strings
	Attr						[]Attribute
	firstChild, NextSiblings	*Node
}

type NodeType int32

const (
	Errornode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val String
}

func parse(r io.REader) (*Node, error)