package main

import (
	"io"
	"io/ioutil"
	"log"
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

func GetNodes(path string, needFiles bool) ([]Node, error) {
	// Reading dir contents.
	files, err := ioutil.ReadDir(path)
	if err != nil {
		println(err.Error())
	}
	var nodes []Node
	for _, f := range files {
		// if we do not need files and current content
		// is not a directory, then we should skip this step.
		if !needFiles && !f.IsDir() {
			continue
		}
		// Node could be just a file.
		node := Node{File: f}

		// However, if it is a directory, it could
		// contain subdirectories.
		if f.IsDir() {
			children, err := GetNodes(path+"/"+f.Name(), needFiles)
			if err != nil {
				println(err.Error())
			}
			node.Children = children
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

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

func dirTree(out io.Writer, path string, printFiles bool) error {
	nodes, err := GetNodes(path, printFiles)
	if err != nil {
		return err
	}
	PrintNodes(out, nodes, "")
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		println("tree [path] [-f]")
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
