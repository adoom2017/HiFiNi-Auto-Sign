package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"hifini/utils"
)

var userAgent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0"

type SignInObject struct {
	URL      string
	Client   *http.Client
	Cookie   string
	message  string
	username string
}

type SignInResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (so *SignInObject) String() string {
	return fmt.Sprintf("用户：%s\n 签到结果: %s\n", so.username, so.message)
}

func (so *SignInObject) Process() error {
	if so.URL == "" {
		so.message = "签到失败"
		return errors.New("url is empty")
	}

	if so.Client == nil {
		so.message = "签到失败"
		return errors.New("client object is null")
	}

	if so.Cookie == "" {
		so.message = "未设置Cookie"
		return errors.New("cookie is empty")
	}

	signPage := so.getSignPage()
	if signPage == "" {
		so.message = "签到失败"
		return errors.New("get sign page failed")
	}

	if strings.Contains(signPage, "请登录") {
		so.message = "Cookie失效"
		return errors.New("invalid cookie")
	}

	sign := utils.GetSign(signPage)
	if sign == "" {
		so.message = "签到失败"
		return errors.New("could not get sign code")
	}

	username := utils.GetUsername(signPage)
	if sign == "" {
		so.message = "签到失败"
		return errors.New("could not get username")
	}

	so.username = username

	msg, err := so.signIn(sign)
	if err != nil {
		so.message = "签到失败"
		return err
	}

	so.message = msg

	return nil
}

func (so *SignInObject) signIn(sign string) (string, error) {
	formData := url.Values{}
	formData.Add("sign", sign)

	log.Println(formData.Encode())

	req, err := http.NewRequest("POST", so.URL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("创建请求失败: ", err)
		return "签到失败", err
	}

	req.Header.Add("Cookie", so.Cookie)
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := so.Client.Do(req)
	if err != nil {
		log.Println("请求签到失败: ", err)
		return "签到失败", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("签到失败")
		return "签到失败", err
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("签到失败: ", err)
		return "签到失败", err
	}

	var result SignInResult
	err = json.Unmarshal(buf, &result)
	if err != nil {
		return "签到失败", err
	}

	return result.Message, nil
}

func (so *SignInObject) getSignPage() string {
	req, err := http.NewRequest("GET", so.URL, nil)
	if err != nil {
		log.Println("创建请求失败: ", err)
		return ""
	}

	req.Header.Add("Cookie", so.Cookie)
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("User-Agent", userAgent)

	//处理返回结果，已经签到过也认为是成功
	resp, err := so.Client.Do(req)
	if err != nil {
		log.Println("请求签到页面失败: ", err)
		return ""
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("请求签到页面失败: ", err)
		return ""
	}
	return string(buf)
}
