package main

import (
	"fmt"
	"io/ioutil"
)

var (
	matches string
)

func fileSearch(root string, filename string) {
	fmt.Println("searching in root", root)
	files, _ := ioutil.ReadDir(root)
}
func main() {
	
	
	
	
	fileSearch("c : /tools", "README.md")
	
}
