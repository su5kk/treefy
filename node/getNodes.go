package node

import "io/ioutil"

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
