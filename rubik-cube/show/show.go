package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"os"
)

//cuda 图片剪切拉伸

var (
	blue      color.Color = color.RGBA{0, 0, 255, 255}
	picwidth  int         = 100
	picheight int         = 100

	W color.Color = color.White
	R color.Color = color.RGBA{255, 0, 0, 255}
	G color.Color = color.RGBA{0, 255, 0, 255}
	O color.Color = color.RGBA{255, 140, 0, 255}
	//OR 88cc ?
	B color.Color = color.RGBA{0, 0, 255, 255}
	Y color.Color = color.RGBA{255, 255, 0, 255}
)

// 大家可以查看这个网址看看这个image包的使用方法 http://golang.org/doc/articles/image_draw.html
func main() {
	//	gen6ColorBlocks()
	add2()

}

func server() {
	http.HandleFunc("/", show)
	http.ListenAndServe("localhost:9999", nil)
}

func show(rw http.ResponseWriter, req *http.Request) {

	//创建一个图像

	m := image.NewRGBA(image.Rect(0, 0, picwidth, picheight)) //*NRGBA (image.Image interface)
	// 填充蓝色,并把其写入到m
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	//以png编码格式,并将m写入到 rw里面去
	png.Encode(rw, m) //Encode writes the Image m to w in PNG format.

	//	png.Encode(os.Stdout, m)

}

func gen6ColorBlocks() {
	outputImage(W, "W")
	outputImage(R, "R")
	outputImage(G, "G")
	outputImage(O, "O")
	outputImage(B, "B")
	outputImage(Y, "Y")
}

func TTT3() {
	//创建一个图像
	m := image.NewRGBA(image.Rect(0, 0, picwidth, picheight)) //*NRGBA (image.Image interface)
	// 填充蓝色,并把其写入到m
	draw.Draw(m, m.Bounds(), &image.Uniform{R}, image.ZP, draw.Src)
	//以png编码格式,并将m写入到 rw里面去
	//	png.Encode(rw, m) //Encode writes the Image m to w in PNG format.
	outputpngFile, err := os.Create("rr" + ".png")
	defer outputpngFile.Close()
	if err == nil {
		png.Encode(outputpngFile, m)
	}
	//	png.Encode(os.Stdout, m)

}

func outputImage(colorc color.Color, filePreName string) {
	//创建一个block 图像
	colorBlock := drawColorBlockImage(colorc)
	//output image to file
	genImageFile(colorBlock, filePreName)
}

//gen png file
func genImageFile(img draw.Image, filePreName string) {
	//以png编码格式,并将m写入到 rw里面去
	//Encode writes the Image m to w in PNG format.
	outputpngFile, err := os.Create(filePreName + ".png")
	defer outputpngFile.Close()
	if err == nil {
		png.Encode(outputpngFile, img)
	}
}

//draw colorblock image
func drawColorBlockImage(colorc color.Color) draw.Image {
	//创建一个black 图像
	black := genBlackPlate()
	m := image.NewRGBA(image.Rect(0, 0, picwidth, picheight)) //*NRGBA (image.Image interface)
	//inner rect
	inRect := image.Rect(10, 10, 90, 90)
	// 填充蓝色,并把其写入到m
	draw.Draw(m, m.Bounds(), &image.Uniform{colorc}, image.ZP, draw.Src)
	//draw colorc block in black background
	draw.Draw(black, inRect, m, image.Pt(0, 0), draw.Src)
	return black
}

func add2() {
	redBlock := drawColorBlockImage(R)
	greenBlock := drawColorBlockImage(G)
	two := addToTop(redBlock, greenBlock)
	genImageFile(two, "addTopTwo")

}

func addToTop(down, up draw.Image) draw.Image {
	m := image.NewRGBA(image.Rect(0, 0, down.Bounds().Size().X, down.Bounds().Size().Y+up.Bounds().Size().Y))
	draw.Draw(m, down.Bounds(), down, image.Pt(0, 0), draw.Src)
	downB := down.Bounds()
	downB2 := downB.Add(image.Pt(0, downB.Max.Y))

	draw.Draw(m, downB2, up, image.Pt(0, 0), draw.Src)

	return m

}

func genBlackPlate() draw.Image {
	m := image.NewRGBA(image.Rect(0, 0, picwidth, picheight)) //*NRGBA (image.Image interface)
	// 填充蓝色,并把其写入到m
	draw.Draw(m, m.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
	return m
}
