package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var inlinebtn1 = tb.InlineButton{
	Text: "🌵",
}

var inlineKeys = [][]tb.InlineButton{
	{inlinebtn1},
}

var urls = []string{
	"https://he8ca29qncg2u39172b39hup-wpengine.netdna-ssl.com/wp-content/uploads/2020/04/rick-astley-never-gonna-give-you-up-meme-696x369.gif",
	//"https://vignette.wikia.nocookie.net/destripando-la-historia/images/4/4d/Zeus-ducha.jpg/revision/latest/scale-to-width-down/200?cb=20200104221621&path-prefix=es",
}

func ShowInlinePics(q *tb.Query) {
	results := make(tb.Results, len(urls))
	for i, url := range urls {
		result := &tb.PhotoResult{
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
}

func SendSticker(m *tb.Message) {
	bot.Send(
		m.Chat,
		"Elegi una opcion",
		&tb.ReplyMarkup{InlineKeyboard: inlineKeys},
	)
}

func GetPicInfo(m *tb.Message) {
	var client = http.Client{}

	pic := tb.File{FileID: m.Photo.FileID}
	file, _ := bot.GetFile(&pic)
	buf1 := new(strings.Builder)
	io.Copy(buf1, file)
	//filebytes, _ := ioutil.ReadAll(file)

	for i := 0; i < len(urls); i++ {
		resp, _ := client.Get(urls[0])
		//body, _ := ioutil.ReadAll(resp.Body)
		buf2 := new(strings.Builder)
		io.Copy(buf2, resp.Body)

		/*if ComparePics(body, filebytes) {
			println("Foto %d OK", i)
		} else {
			println("Foto %d NOT OK", i)
		}*/
		resp.Body.Close()
	}
}

/*
func ComparePics(p1, p2 []byte) bool {
	var result = true
	if len(p1) != len(p2) {
		println("%d - %d", len(p1), len(p2))
		return false
	}
	for i, val := range p1 {
		if val != p2[i] {
			result = false
			break
		}
	}
	return result
}*/
