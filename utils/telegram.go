package utils

import (
	"fmt"
	"net/http"
	"net/url"
)

func SendMessage(botToken, chatID, text string) {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	data := url.Values{
		"chat_id": {chatID},
		"text":    {text},
	}

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		fmt.Printf("Error sending message: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Message sent, response status: %s\n", resp.Status)
}
