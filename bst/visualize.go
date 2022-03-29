package bst

import (
	"fmt"
	"os"
	"os/exec"
)

//Visualize creates a dotty file and runs the dot program (please make sure to have GraphViz initialized)
func (b BinaryTree[T]) VisualizeDotty(fileName string) {
	file, err := os.Create(fileName + ".dot")

	defer file.Close()

	if err == nil {

		_, err2 := file.WriteString("digraph BinarySearchTree{\n")
		if err2 != nil {fmt.Println(err2); return }

		recurseWrite(file, b.root)

		_, err3 := file.WriteString("}")
		if err3 != nil {fmt.Println(err3); return}

		out, err4 := exec.Command("dot", "-Tjpg ./" + fileName + ".dot -o " + fileName + ".jpg").Output()
		if err4 != nil {
			
			output := string(out[:])
			fmt.Println("Dot Error:", output, ";", err4)
		}

		return
	}

	fmt.Println(err)
	//string command = "dot -Tjpg " + outputFilename + " -o " + jpgFilename;
	//system(command.c_str());
}

func recurseWrite[T any](file *os.File, node *binaryNode[T]){
	
	_, err1 := file.WriteString(fmt.Sprint(node.id) + " [label=\"" + fmt.Sprintf("%v", node.value) + "\"]\n")
	if err1 != nil { fmt.Println(err1); return}

	if node.right != nil {
		_, err2 := file.WriteString(fmt.Sprint(node.id) + " -> " + fmt.Sprint(node.right.id) + "\n")
		if err2 != nil { fmt.Println(err1); return}

		recurseWrite(file, node.right)
	}

	if node.left != nil {
		_, err3 := file.WriteString(fmt.Sprint(node.id) + " -> " + fmt.Sprint(node.left.id) + "\n")
		if err3 != nil { fmt.Println(err1); return}

		recurseWrite(file, node.left)
	}
}