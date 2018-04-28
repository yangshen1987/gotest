package html

import "io"

type Node struct {
	Type NodeType
	Data string
	Attr []Arribute
	FristChild, NextSibling *Node
}
type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	CommentNode
	DoctypeNode
)

type Arribute struct {
	Key, Val string
}

