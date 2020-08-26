package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strconv"

	//	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	//"regexp"
	// "strings"
)

func main() {
	fmt.Println("aa")
	//ffmpeg -i audio.m4s -i video.m4s -codec copy test.mp4
	//dir := `E:\bilibili\tv.danmaku.bili\download\200700409`
	dir := `/mnt/e/bilibili/tv.danmaku.bili/download/200700409`

	file, err := os.Open(dir)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	dir_list, err := file.Readdirnames(0) // 0 to read all files and folders
	if err != nil {
		log.Fatalf("failed reading directory: %s", err)
	}
	for _, subdir := range dir_list {
		jsfile := filepath.Join(dir, subdir, "entry.json")
		//fmt.Println(jsfile)
		jsfile2, err := os.Open(jsfile)

		if err != nil {
			panic(err)
		}
		content, err1 := ioutil.ReadAll(jsfile2)
		if err1 != nil {
			panic(err1)
		}
		// jsonStr := string(content)
		//fmt.Println(jsonStr)
		var v interface{}
		json.Unmarshal(content, &v)
		v2 := v.(map[string]interface{})
		pageData := v2["page_data"].(map[string]interface{})

		mp4fileName := num2str(pageData["page"].(float64)) + pageData["part"].(string) + ".mp4"
		//fmt.Println(mp4fileName)
		//ffmpeg -i audio.m4s -i video.m4s -codec copy test.mp4
		// fmt.Println("ffmpeg", "-i", filepath.Join(dir, subdir, "audio.m4s"), "-i", filepath.Join(dir, subdir, "video.m4s"), "-codec", "copy",
		// 	filepath.Join(dir, mp4fileName))
		// cmd := exec.Command("echo", mp4fileName)

		cmdArguments := []string{"-i", filepath.Join(dir, subdir, "audio.m4s"), "-i", filepath.Join(dir, subdir, "video.m4s"),
			"-codec", "copy", filepath.Join(dir, mp4fileName)}

		cmd := exec.Command("ffmpeg", cmdArguments...)
		// cmd := exec.Command("ffmpeg", " -i "++" -i "++
		// 	" -codec  copy "+)
		// cmd.Run()
		//cmd.Run()

		var out bytes.Buffer

		cmd.Stdout = &out
		// err2 := cmd.Run()
		// if err2 != nil {
		// 	log.Fatal(err2)
		// }
		fmt.Printf("command output: %q", out.String())

		// str, err := cmd.Output()

		// fmt.Println(str, err)
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
		// fmt.Println(string(str))
		//fmt.Println(mp4fileName)

		jsfile2.Close()
		// ./1/80/audio.m4s video.m4s  ./1/entry.json
		// oldName := name
		// fmt.Println("Old Name - ", oldName)
		// newName := strings.Split(oldName, "：")[1]
		// fmt.Println("New Name - ", newName)
		//err := os.Rename(filepath.Join(dir, oldName), filepath.Join(dir, newName))
		// if err != nil {
		// 	continue
		// }
		//fmt.Println("File names have been changed")
	}

}

func num2str(a float64) string {
	if (a) < 10 {

		return "00" + strconv.Itoa(int(a))

	} else if a < 100 {
		return "0" + strconv.Itoa(int(a))
	} else {
		return strconv.Itoa(int(a))
	}
}
