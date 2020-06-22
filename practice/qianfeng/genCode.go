package main

import (
	"fmt"
	// "io"
	// "io/ioutil"
	"os"
	"strconv"
)

func main() {
	for i := 1; i < 180; i++ {
		if i != 115 && i != 116 {
			dirName := "./p" + strconv.Itoa(i)
			if !FileExist(dirName) {
				err := os.Mkdir(dirName, os.ModePerm)
				if err != nil {
					fmt.Println(err)
					return
				}
				goFile, err := os.Create(dirName + "/" + dirName + ".go")
				if err != nil {
					fmt.Println(err)
					return
				}
				str := `package main

import (
	"fmt"
)

func main() {
	
}`
				goFile.WriteString(str)
				goFile.Close()
				// ioutil.WriteFile(goFile)
			}

		}

	}
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
