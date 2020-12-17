package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main(){
	//1、批量生成随机txt文件，2、开n个协程并行上传文件，记录结果
	fmt.Println("usage:",os.Args[0],"ipfsUrl(127.0.0.1:5001/api/v0/add) fileNum(2000) parallelNum(50) urlParam(chunker=size-1024&quieter=true")
	//单次测试文件总数
	fileNum := 2000
	//并发数
	parallelNum:=50
	fileNamePre := "bb"
	ipfsUrl:="http://127.0.0.1:5001/api/v0/add"
	urlParam:="chunker=size-1024&quieter=true"
	l:=len(os.Args)

	if l>1 {
		ipfsUrl=os.Args[1]
	}
	if l>2 {
		i,_:=strconv.Atoi(os.Args[2])
		fileNum = i
	}
	if l>3 {
		i,_:=strconv.Atoi(os.Args[3])
		parallelNum = i
	}

	if l>4 {
		urlParam=os.Args[4]
	}
	fileDir:="files"
	os.Mkdir(fileDir,0755)
	GenFile(fileNamePre,1000,fileNum)
	fmt.Println("test start")
	t1:=time.Now()
	wg.Add(fileNum)
	ch := make(chan string,500)
	ch2:=make(chan time.Duration,fileNum)
	for i:=0;i<parallelNum;i++{
		go ParallelRequest(ipfsUrl,urlParam,fileDir,ch,ch2)
	}
	wg.Wait()
	fmt.Println("test end")
	var sumMs time.Duration =0
	for i:=0;i<fileNum;i++{
		t := <-ch2
		sumMs+=t
	}
	avg := sumMs.Milliseconds()/int64(fileNum)
	fmt.Println("http cost avg ms:",avg)
	totalCost := time.Since(t1)
	resultFile,_:=os.OpenFile("result.txt",os.O_RDWR | os.O_CREATE | os.O_APPEND,0644)
	result := fmt.Sprintf("tps:%v,threads:%v,files:%v,avgMs:%v,time:%v,url:%s,urlParam:%s",float64(fileNum)/totalCost.Seconds(),)
	fmt.Println(result)
	resultFile.Write([]byte(result))
}

func ParallelRequest(ipfsUrl string, param string, fileDir string, ch chan string, ch2 chan time.Duration) {
	for {
		fileName := <-ch
		t1 := time.Now()
		RequestIPFS(ipfsUrl,param,fileDir,fileName,ch2)
		cost := time.Since(t1)
		ch2 <- cost
	}

}

func RequestIPFS(url string, param string, dir string, fileName string, ch2 chan time.Duration) {
	b,w:=CreateMultipartFormData("file",dir,fileName)
	req,err := http.NewRequest("POST",url+"?"+param,&b)
	if err !=nil {
		fmt.Println("request err:",err)
	}
	req.Header.Set("Content-Type",w.FormDataContentType())
	c:=&http.Client{}
	response, _ := c.Do(req)
	ioutil.ReadAll(response.Body)
	response.Body.Close()
	wg.Done()
}

func GenFile(fileNamePre string, fileLength int, num int) {
 rand.Seed(time.Now().UnixNano())
 for i:=0;i<num; i++ {
 	str:=GenRandomString(fileLength)
 	ioutil.WriteFile(fmt.Sprintf("./files/%v%v.txt",fileNamePre,i),[]byte(str),os.ModePerm)
 }
}

func GenRandomString(length int)  string {
	var letters =[]rune("abcd")
	s:=make([]rune,length)
	for i:=range s {
		s[i] = letters[rand.Intn(len(letters))]

	}
	return string(s)

}

func CreateMultipartFormData(fieldName,dir,fileName string)( bytes.Buffer,*multipart.Writer){
	var b bytes.Buffer
	var err error
	w:= multipart.NewWriter(&b)
	var fw io.Writer
	file := mustOpen(dir +"/" + fileName)

	if fw,err = w.CreateFormFile(fieldName,fileName);err!=nil{
		fmt.Println("err create writer",err)
	}
	if _,err = io.Copy(fw,file);err != nil {
		fmt.Println("copy err:",err)
	}
	file.Close()
	w.Close()
	return b,w
}

func mustOpen(fileName string) *os.File {
	f,err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	return f

}