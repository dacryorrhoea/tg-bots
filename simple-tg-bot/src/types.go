package main

// container for bot
type Bot struct {
	Token  string
	Api    string
	Url    string
	Offset int
}

func (bot *Bot) init(token string, api string) {
	bot.Token = token
	bot.Api = api
	bot.Url = api + token
	bot.Offset = 0
}

// tg api
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}
