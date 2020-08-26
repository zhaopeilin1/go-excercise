package main

import (
	//	"io/ioutil"
	"fmt"
	"log"
	"os"
	"path/filepath"

	//"regexp"
	"strings"
)

func main() {
	dir := `E:\bilibili\tv.danmaku.bili\download\go`

	file, err := os.Open(dir)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, err := file.Readdirnames(0) // 0 to read all files and folders
	if err != nil {
		log.Fatalf("failed reading directory: %s", err)
	}
	for _, name := range list {
		oldName := name
		fmt.Println("Old Name - ", oldName)
		newName := strings.Split(oldName, "：")[1] 
		fmt.Println("New Name - ", newName)
		err := os.Rename(filepath.Join(dir, oldName), filepath.Join(dir, newName))
		if err != nil {
			continue
		}
		fmt.Println("File names have been changed")
	}
}
