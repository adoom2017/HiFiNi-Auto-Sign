package main

import (
	"hifini/model"
	"hifini/utils"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	so := model.SignInObject{
		URL:    "https://www.hifini.com/sg_sign.htm",
		Client: &http.Client{Timeout: time.Second * 3},
		Cookie: os.Getenv("COOKIES"),
	}

	err := so.Process()
	if err != nil {
		log.Println("签到失败：", err)
	} else {
		log.Println("签到成功：", so.String())
	}

	token := os.Getenv("TG_TOKEN")
	id := os.Getenv("TG_CHAT_ID")
	if token != "" && id != "" {
		t := utils.Telegram{
			Token:  token,
			ChatID: id,
		}

		utils.Sending(t, so.String())
	}
}
