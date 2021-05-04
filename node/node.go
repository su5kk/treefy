package node

import (
	"os"
	"strconv"
)

type Node struct {
	File     os.FileInfo
	Children []Node
}

func (n *Node) Size() string {
	if n.File.Size() > 0 {
		return strconv.FormatInt(n.File.Size(), 10) + "b"
		// return fmt.Sprintf("%vb", n.File.Size())
	}
	return "0b"
}

// If the node is a directory, return name
// else return name and size
func (n *Node) Name() string {
	if n.File.IsDir() {
		return n.File.Name()
	}
	return n.File.Name() + " (" + n.Size() + ")"
	// Omitted due to performance constraints. Used concatenation instead.
	// return fmt.Sprintf("%v (%v)", n.File.Name(), n.Size())
}
