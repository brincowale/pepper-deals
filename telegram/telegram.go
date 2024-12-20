package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"pepper-deals/pepper"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	Request *http.Client
	Token   string
}

type Message struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

type TelegramResponse struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func New(token string) Client {
	httpClient := &http.Client{
		Timeout: time.Second * 30,
	}
	return Client{
		Request: httpClient,
		Token:   token,
	}
}

func (c Client) SendMessage(chatId string, message string) error {
	data, err := json.Marshal(
		Message{
			ChatId: chatId,
			Text:   message,
		},
	)
	if err != nil {
		return err
	}
	endpoint := "https://api.telegram.org/bot" + c.Token + "/sendMessage?parse_mode=HTML"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Request.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response TelegramResponse
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return err
	}
	if response.Ok {
		return nil
	}
	strError := strconv.Itoa(response.ErrorCode) + ": " + response.Description
	return errors.New(strError)
}

func (c Client) CreateMessage(deal pepper.Deal) string {
	returnLine := "\n\n"
	var price string
	if deal.Price != 0.0 {
		price = "Price: " + strconv.FormatFloat(deal.Price, 'f', 2, 64) + "€" + returnLine
	}
	var categories string
	for _, category := range deal.Groups {
		categories += category.Name + " | "
	}
	deal.Merchant.URLName = strings.ReplaceAll(deal.Merchant.URLName, "-", ".")
	str := deal.Title + returnLine + stripHTMLTags(deal.Description) + returnLine + categories + returnLine + price + deal.DealURI + returnLine + deal.Merchant.URLName
	return str
}

func stripHTMLTags(input string) string {
	var output string
	inTag := false
	for _, char := range input {
		if char == '<' {
			inTag = true
		} else if char == '>' {
			inTag = false
			output += " "
		} else if !inTag {
			output += string(char)
		}
	}
	return output
}
