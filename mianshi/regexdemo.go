package main

import (
	"fmt"
	"regexp"
)

func main() {

	var chapTitleRegEx = "^.{0,8}[第]? *[一二三四五六七八九十百千万两1234567890 ]+[章]"
	re := regexp.MustCompile(chapTitleRegEx)

	str1 := "第001章 重生"
	str2 := "贫民区的一间低矮小屋门前，陆离拍了拍小萝莉的脑袋，爱怜的揉了又揉，能看到妹妹如此健康活泼的活着，真好！"

	fmt.Println(re.MatchString(str1), re.MatchString(str2))
}
