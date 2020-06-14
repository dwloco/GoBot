package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

var bot *tb.Bot

func main() {
	token := GetToken()
	if token != nil {
		b, err := tb.NewBot(tb.Settings{
			Token:  token.(string),
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		bot = b
		SetHandlers()
		bot.Start()
	}
}

func SetHandlers() {
	bot.Handle(tb.OnQuery, ShowInlinePics)
	bot.Handle(tb.OnText, UpdateGlobalGroupID)
	bot.Handle("/hola", SayHi)
	bot.Handle("/sendsticker", SendSticker)
	bot.Handle(tb.OnChosenInlineResult, GetPicInfo)
}
