package main

import (
	"fmt"
	"os"
)

func main() {

	listDir("E:\\software\\GoWorks\\src\\xuexiangban_go\\lesson11", 0)

}

func listDir(filePath string, tabint int) {

	tab := "|--"

	for i := 0; i < tabint; i++ {
		tab = "	 " + tab
	}

	dir := filePath

	readDir, _ := os.ReadDir(dir)
	for _, file := range readDir {
		fileName := file.Name()
		fmt.Printf("%s%s\n", tab, fileName)
		if file.IsDir() {
			listDir(fileName, tabint+1)
		}
	}
}
