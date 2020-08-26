package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	//TODO 声音，水声
	time.Sleep(4 * time.Second)
	x, y := seekTag(554, 99, 555, 355, "192f39")
	fmt.Println(x, y)
	// screen()
	// getMousePositionAndColor()
	//pos:  554 99 左上
	// color----  203742

	// pos:  929 355 右下
	// color----  14292f

	//pos:  780 205
	// color----  192f39 鱼漂颜色

}

func moveMose() {
	//TODO 鼠标形状
	robotgo.ScrollMouse(10, "up")
	robotgo.MouseClick("left", true)
	// robotgo.MouseToggle()
	robotgo.GetMouseColor()
	robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
}

func keyboard() {

	//TODO
	robotgo.TypeStr("Hello World")
	robotgo.TypeStr("だんしゃり", 1.0)
	// robotgo.TypeString("テストする")

	robotgo.TypeStr("Hi galaxy. こんにちは世界.")
	robotgo.Sleep(1)

	// ustr := uint32(robotgo.CharCodeAt("Test", 0))
	// robotgo.UnicodeType(ustr)

	robotgo.KeyTap("enter")
	// robotgo.TypeString("en")
	robotgo.KeyTap("i", "alt", "command")

	arr := []string{"alt", "command"}
	robotgo.KeyTap("i", arr)

	robotgo.WriteAll("Test")
	text, err := robotgo.ReadAll()
	if err == nil {
		fmt.Println(text)
	}
}

func getMousePositionAndColor() {
	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)
	color := robotgo.GetPixelColor(x, y)
	fmt.Println("color---- ", color)
}

//从x1,y1到x2,y2范围内查找颜色
func seekTag(x1, y1, x2, y2 int, color string) (int, int) {
	t1 := time.Now()
	fmt.Println((x1 - x2) * (y1 - y2))
	for i := y1; i <= y2; i = i + 5 {
		for j := x1; j <= x2; j = j + 5 {
			// color :=
			robotgo.GetPixelColor(i, j)
			// fmt.Println(color)
		}
	}
	fmt.Println((x1-x2)*(y1-y2), ":cost:", time.Since(t1))
	return 0, 0

}

func screen() {

	x, y := robotgo.GetMousePos()
	fmt.Println("pos: ", x, y)

	color := robotgo.GetPixelColor(100, 200)
	fmt.Println("color---- ", color)

	color2 := robotgo.GetPixelColor(x, y)
	fmt.Println("color---- ", color2)

}
