package main

import (
	"encoding/json"
	"fmt"
	tb "gopkg.in/tucnak/telebot.v3"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var inlinebtn1 = tb.InlineButton{
	Unique:          "Cacti",
	Text:            "🌵",
	InlineQuery: "Cacti",
}

var inlinebtn2 = tb.InlineButton{
	Unique:          "Turtle",
	Text:            "🐢",
}

var inlineKeys = [][]tb.InlineButton{
	{inlinebtn1},
}


var bot *tb.Bot

func main() {
	token := GetToken()
	if token != nil {
		b, err := tb.NewBot(tb.Settings{
			Token: token.(string),
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
	bot.Handle("/hola", func(m *tb.Message) {
		bot.Send(m.Chat, "Holis")
	})
	bot.Handle("/sendsticker", func(m *tb.Message) {
		bot.Send(
			m.Chat,
			"Elegi una opcion",
			&tb.ReplyMarkup{InlineKeyboard: inlineKeys},
		)
	})
	bot.Handle(&inlinebtn1, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
		})
		var result tb.ArticleResult
		result.SetContent(&tb.InputTextMessageContent{
			Text: "Some text",
		})
		bot.Send(c.Sender, result)
		//bot.Send(c.Message.Chat, "Puto el que lee 1")
	})
	/*bot.Handle(&inlinebtn2, func(c *tb.Callback) {
		bot.Respond(c, &tb.CallbackResponse{
			ShowAlert: false,
			Text: "xd",
		})
		//bot.Send(c.Message.Chat, "Puto el que lee 2")
	})*/
}

func GetToken() interface{} {
	file, err := os.Open("token.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	val, _ := ioutil.ReadAll(file)
	var result map[string]interface{}
	json.Unmarshal(val, &result)
	return result["token"]
}


func ShowInlinePics(q *tb.Query) {
	urls := []string {
		"https://he8ca29qncg2u39172b39hup-wpengine.netdna-ssl.com/wp-content/uploads/2020/04/rick-astley-never-gonna-give-you-up-meme-696x369.gif",
		"https://vignette.wikia.nocookie.net/destripando-la-historia/images/4/4d/Zeus-ducha.jpg/revision/latest/scale-to-width-down/200?cb=20200104221621&path-prefix=es",
	}
	results := make(tb.Results, len(urls))
	for i, url := range urls {
		result := &tb.PhotoResult{
			URL: url,
			ThumbURL: url,
		}
		results[i] = result
		results[i].SetResultID(strconv.Itoa(i))
	}
	err := bot.Answer(q, &tb.QueryResponse{
		Results: results,
		CacheTime: 60,
	})
	if err != nil {
		println(err)
	}
}