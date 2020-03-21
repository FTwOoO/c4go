package ast

import (
	"testing"
)

func TestColdAttr(t *testing.T) {
	nodes := map[string]Node{
		`0x7fc9bf0d51d8 <col:33>`: &ColdAttr{
			Addr:       0x7fc9bf0d51d8,
			Pos:        NewPositionFromString("col:33"),
			ChildNodes: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
