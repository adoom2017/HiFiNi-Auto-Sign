package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	client := &http.Client{}

	for i := 0; i < 2; i++ {
		result := SignIn(client)
		if strings.Contains(result, "成功") || strings.Contains(result, "已经签过") {
			fmt.Println("签到成功!!!")
			return
		} else if strings.Contains(result, "正在进行人机识别") {
			fmt.Println("人机识别，等待后重试")
			time.Sleep(5 * time.Second)
			continue
		} else {
			fmt.Println("签到失败!!!")
			os.Exit(3)
		}
	}
}

// SignIn 签到
func SignIn(client *http.Client) string {
	//生成要访问的url
	url := "https://www.hifini.com/sg_sign.htm"
	cookie := os.Getenv("COOKIES")
	if cookie == "" {
		fmt.Println("COOKIES不存在，请检查是否添加")
		return ""
	}
	//提交请求，修改变化
	reqest, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("创建请求失败: ", err)
		return ""
	}

	reqest.Header.Add("Cookie", cookie)
	reqest.Header.Add("x-requested-with", "XMLHttpRequest")
	//处理返回结果，已经签到过也认为是成功
	response, err := client.Do(reqest)
	if err != nil {
		fmt.Println("请求签到失败: ", err)
		return ""
	}
	defer response.Body.Close()
	fmt.Println("响应Cookies:", response.Header.Get("Cookie"))

	buf, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("获取签到结果失败: ", err)
		return ""
	}

	fmt.Println(string(buf))
	return string(buf)
}
