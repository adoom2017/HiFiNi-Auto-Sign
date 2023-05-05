package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}
	success := SignIn(client)
	if success {
		fmt.Println("签到成功!")
	} else {
		fmt.Println("签到失败!")
		os.Exit(3)
	}
}

// SignIn 签到
func SignIn(client *http.Client) bool {
	//生成要访问的url
	url := "https://www.hifini.com/sg_sign.htm"
	cookie := os.Getenv("COOKIES")
	if cookie == "" {
		fmt.Println("COOKIES不存在，请检查是否添加")
		return false
	}
	//提交请求
	reqest, err := http.NewRequest("POST", url, nil)
	reqest.Header.Add("Cookie", cookie)
	reqest.Header.Add("x-requested-with", "XMLHttpRequest")
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	buf, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(buf))
	return strings.Contains(string(buf), "成功")
}
