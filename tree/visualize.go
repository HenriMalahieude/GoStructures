package tree

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

//Visualize creates a dotty file and runs the dot program (please make sure to have GraphViz initialized). Note: Dotty mirrors the tree for some reason
func (b BinaryTree[T]) VisualizeDotty(fileName string) {
	file, err := os.Create(fileName + ".dot")

	if err == nil {
		defer file.Close()

		_, err2 := file.WriteString("digraph BinarySearchTree{\n")
		if err2 != nil {fmt.Println(err2); return }

		recurseWrite(file, b.root)

		_, err3 := file.WriteString("}")
		if err3 != nil {fmt.Println(err3); return}

		cmd := exec.Command("dot", "-Tjpg", "./" + fileName + ".dot", "-o", fileName + ".jpg")
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		err4 := cmd.Run()
		if err4 != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		}

		return
	}

	fmt.Println(err)
	//string command = "dot -Tjpg " + outputFilename + " -o " + jpgFilename;
	//system(command.c_str());
}

func recurseWrite[T any](file *os.File, node *BinaryNode[T]){
	
	_, err1 := file.WriteString(fmt.Sprint(node.id) + " [label=\"" + fmt.Sprintf("%v", node.value) + "\"]\n")
	if err1 != nil { fmt.Println(err1); return}

	if node.right != nil && node.right != node {
		_, err2 := file.WriteString(fmt.Sprint(node.id) + " -> " + fmt.Sprint(node.right.id) + "\n")
		if err2 != nil { fmt.Println(err1); return}

		recurseWrite(file, node.right)
	}

	if node.left != nil && node.left != node {
		_, err3 := file.WriteString(fmt.Sprint(node.id) + " -> " + fmt.Sprint(node.left.id) + "\n")
		if err3 != nil { fmt.Println(err1); return}

		recurseWrite(file, node.left)
	}
}