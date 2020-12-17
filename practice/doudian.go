package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	//"github.com/RivenZoo/blevejieba"
)

func main() {
	//p := path.Join("/aa", "bb.txt")
	//fmt.Println(p)
	//
	//secret := "7b7eb452-5741-45e7-83a0-2b2a7847e2c8"
	//var myStr string = secret + "app_key6854476287269373448" + "methodproduct.getGoodsCategory" + "param_json{\"cid\":\"0\"}" + "timestamp2020-07-29 17:30:36" + "v2" + secret
	//fmt.Println(myStr)
	//str := getMd5String1(myStr)
	//fmt.Println(str)
	//fmt.Println(getMd5String2(myStr))
}

//func jiebademo(){
//	blevejieba
//
//}

func getMd5String2(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		log.Fatal(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}
func getMd5String1(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		log.Fatal(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}
