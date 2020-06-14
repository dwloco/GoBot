package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

var inlinebtn1 = tb.InlineButton{
	Text: "🌵",
}

var inlineKeys = [][]tb.InlineButton{
	{inlinebtn1},
}

var urls = []string{
	"https://he8ca29qncg2u39172b39hup-wpengine.netdna-ssl.com/wp-content/uploads/2020/04/rick-astley-never-gonna-give-you-up-meme-696x369.gif",
	"https://vignette.wikia.nocookie.net/destripando-la-historia/images/4/4d/Zeus-ducha.jpg/revision/latest/scale-to-width-down/200?cb=20200104221621&path-prefix=es",
}

const (
	RickRoll = "0"
	Zeus     = "1"
)

var messageQueue []string

var group *tb.Chat

func ShowInlinePics(q *tb.Query) {
	results := make(tb.Results, len(urls))
	for i, url := range urls {
		result := &tb.PhotoResult{
			//Description: strconv.Itoa(i),
			URL:      url,
			ThumbURL: url,
		}
		results[i] = result
		results[i].SetResultID(strconv.Itoa(i))
	}
	err := bot.Answer(q, &tb.QueryResponse{
		Results:   results,
		CacheTime: 60,
	})
	if err != nil {
		println(err)
	}
}

func SayHi(m *tb.Message) {
	bot.Send(m.Chat, "Holis")
	group = m.Chat
}

func SendSticker(m *tb.Message) {
	bot.Send(
		m.Chat,
		"Elegi una opcion",
		&tb.ReplyMarkup{InlineKeyboard: inlineKeys},
	)
}

func GetPicInfo(c *tb.ChosenInlineResult) {
	if group != nil {
		AnswerPic(c.ResultID)
	} else {
		messageQueue = append(messageQueue, c.ResultID)
	}
}

func AnswerPic(id string) {
	switch id {
	case RickRoll:
		bot.Send(group, "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	case Zeus:
		bot.Send(group, "ZEUS")
	default:
		fmt.Println("Codigo no definido (" + id + ")")
	}
}

func UpdateGlobalGroupID(m *tb.Message) {
	//println(m.Chat.ID)
	group = m.Chat
	if messageQueue != nil {
		for _, val := range messageQueue {
			AnswerPic(val)
		}
		messageQueue = nil
	}
}
