package ast

import (
	"strings"

	"github.com/FTwOoO/c4go/util"
)

// AllocSizeAttr is a type of attribute that is optionally attached to a variable
// or struct field definition.
type AllocSizeAttr struct {
	Addr        Address
	Pos         Position
	IsInherited bool
	A           int
	B           int
	ChildNodes  []Node
}

func parseAllocSizeAttr(line string) *AllocSizeAttr {
	groups := groupsFromRegex(
		`<(?P<position>.*)>(?P<inherited> Inherited)?(?P<a> \d+)(?P<b> \d+)?`,
		line,
	)

	var av, bv int
	a := strings.TrimSpace(groups["a"])
	if a != "" {
		av = util.Atoi(a)
	}
	b := strings.TrimSpace(groups["b"])
	if b != "" {
		bv = util.Atoi(b)
	}

	return &AllocSizeAttr{
		Addr:        ParseAddress(groups["address"]),
		Pos:         NewPositionFromString(groups["position"]),
		IsInherited: len(groups["inherited"]) > 0,
		A:           av,
		B:           bv,
		ChildNodes:  []Node{},
	}
}

// AddChild adds a new child node. Child nodes can then be accessed with the
// Children attribute.
func (n *AllocSizeAttr) AddChild(node Node) {
	n.ChildNodes = append(n.ChildNodes, node)
}

// Address returns the numeric address of the node. See the documentation for
// the Address type for more information.
func (n *AllocSizeAttr) Address() Address {
	return n.Addr
}

// Children returns the child nodes. If this node does not have any children or
// this node does not support children it will always return an empty slice.
func (n *AllocSizeAttr) Children() []Node {
	return n.ChildNodes
}

// Position returns the position in the original source code.
func (n *AllocSizeAttr) Position() Position {
	return n.Pos
}
