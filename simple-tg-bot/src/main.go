package main

import (
	"log"
)

func main() {
	bot := new(Bot)
	bot.init("2043706009:AAGyOD6rkIbtv93jtWFBeEgCAoBoXYlP98I", "https://api.telegram.org/bot")

	for {
		updateList, err := request(bot)
		if err != nil {
			log.Println("Smth went wrong: ", err.Error())
		}

		for _, update := range updateList {
			botMessage := BotMessage{
				update.Message.Chat.ChatId,
				"",
			}

			switch update.Message.Text {
			case "привет":
				botMessage.Text = "привет"
			case "пока":
				botMessage.Text = "удачи"
			default:
				botMessage.Text = "?"
			}

			err = respond(bot, botMessage)
			if err != nil {
				log.Println("Smth went wrong: ", err.Error())
			}

			bot.Offset = update.UpdateId + 1
		}
	}
}
