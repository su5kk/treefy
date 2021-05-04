package main

import (
	"io"
	"log"
	"os"

	node "github.com/su5kk/treefy/node"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	nodes, err := node.GetNodes(path, printFiles)
	if err != nil {
		return err
	}
	node.PrintNodes(out, nodes, "")
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		println("treefy [path] [-f]")
		return
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
