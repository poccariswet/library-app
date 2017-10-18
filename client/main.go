package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Isbn13 string `json:"isbn13"`
	State  bool   `json:"state"`
}

const url = "http://localhost:9090/book/pullbookinfo/"

var book_state string

func get(url string, info interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, info)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					api_url := url + message.Text
					fmt.Println(api_url)
					book := Book{}
					err := get(api_url, &book)

					if book.Id == 0 {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("そのタイトルの本はありませんでした。")).Do(); err != nil {
							log.Print(err)
						}
						continue
					}
					if book.State {
						book_state = message.Text + "は貸し出し可能です"
					} else {
						book_state = message.Text + "は貸出中のため貸し出しできません"
					}

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(book_state)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
