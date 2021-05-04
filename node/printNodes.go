package node

import (
	"io"
	"log"
)

func PrintNodes(out io.Writer, nodes []Node, parentPrefix string) {
	prefix := "├───"
	childPrefix := "│\t"

	for i, node := range nodes {
		if i == len(nodes)-1 {
			prefix = "└───"
			childPrefix = "\t"
		}
		_, err := out.Write([]byte(parentPrefix + prefix + node.Name() + "\n"))
		if err != nil {
			log.Fatal(err)
		}

		// Omitted due to performance issues. Used io.Writer.Write instead.
		//fmt.Fprint(out, parentPrefix, prefix, node.Name(), "\n")

		if node.File.IsDir() {
			PrintNodes(out, node.Children, parentPrefix+childPrefix)
		}
	}
}
