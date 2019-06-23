package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

//实例化复合类型，使用大括号,颜色数组切片。作为调色板
var palette = []color.Color{color.White, color.RGBA{0x00, 0xFF, 0x00, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	//实例化复合类型，使用大括号
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	//外部循环，一次循环画出动画的一帧
	//每一次都会生成一个单独的动画
	//帧。它生成了一个包含两种颜色的201*201大小的图片,白色和黑色。所有像素点都会被默认设置
	//为其零值(也就是调色板palette里的第0个值),这里我们设置的是白色
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		//使用颜色数组调色板，生成一个调色图片
		img := image.NewPaletted(rect, palette)
		//内部循环，角度从0到2Pi。完成一幅
		//内层循环设置两个偏振值。x轴偏振使用sin函数。y轴偏振也是正弦波,但其相对x轴的偏振是一个
		//0­3的随机值,初始偏振值是一个零值,随着动画的每一帧逐渐增加。循环会一直跑到x轴完成五次
		//完整的循环。每一步它都会调用SetColorIndex来为(x, y)点来染黑色。
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
