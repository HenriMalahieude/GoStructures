package main

import (
	"fmt"
	"runtime/debug"
)

func catch(origin string){
	if r := recover(); r != nil{
		fmt.Println("Recovered an error in " + origin + ":", r)
		fmt.Println("Stacktrace: \n" + string(debug.Stack()))
	}	
}

func doCheck(tNum *int, out string, ff func() bool){
	if ff() {
		(*tNum)++;
	}else{
		fmt.Println(out)
	}
}

func testSingleList() {

}

func testDoubleList() {

}

func testCircleList() {

}

func main() {
	testStack()
	testQueue()
	testBST()
}