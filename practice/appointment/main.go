package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Jeffail/gabs/v2"
	// "github.com/Jeffail/gabs"

	//"fmt"
	//"io"
	"net/http"
	//	"os"
	"strings"
)

func main() {
	queryUrl := "https://gzhzyw.gzjd.gov.cn/api/dtgl/web/serviceEntry?ssfwdt=GZGA_ZCXT"
	queryData := `{"Service":"Queue.GetDateCountByFwdtAndItem","yyFwdtCode":"GZGA_ZCXT","yyItemCode":"440100-172-FW-027-01","yyOrderCycle":16}`

	resp, _ := http.Post(queryUrl, "application/json", strings.NewReader(queryData))
	body := resp.Body
	defer body.Close()

	content, err := ioutil.ReadAll(body)

	if err != nil {
		fmt.Errorf("read error:", err)
	}
	fmt.Println(string(content))

	jsonParsed, err := gabs.ParseJSON(content)
	//YYDateCount  jsonParsed.Path("outter.inner.value1").Data().(float64)

	for _, child := range jsonParsed.S("YYDateCount").Children() {
		date := child.Path("yy_date")
		if date.Data().(string) == "2020-10-01" {
			count := child.S("leftYYCount").Data().(float64)
			if count > 0 {
				http.Get("https://sc.ftqq.com/SCU111982T529d51677fcf143188eb4eb65fa33c285f5054d471813.send?text=有空位，快预约" + fmt.Sprint(count))
			} else {
				//http.Get("https://sc.ftqq.com/SCU111982T529d51677fcf143188eb4eb65fa33c285f5054d471813.send?text=无空位" + fmt.Sprint(count))
			}
		}

	}
	if err != nil {
		panic(err)
	}

}
