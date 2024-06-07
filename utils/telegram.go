package utils

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Telegram struct {
	Token  string
	ChatID string
}

func (t Telegram) SendMessage(text string) {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.Token)

	data := url.Values{
		"chat_id": {t.ChatID},
		"text":    {text},
	}

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		log.Printf("Error sending message: %v\n", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("Message sent, response status: %s\n", resp.Status)
}
