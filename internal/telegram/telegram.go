package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	tgbotapi "github.com/msiddikov/telegram-bot-api/v6"
)

var (
	bot *tgbotapi.BotAPI
)

func StartTelegram() {
	godotenv.Load(".env")
	token := os.Getenv("TG_TOKEN")
	bot, _ = tgbotapi.NewBotAPI(token)

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			msg := update.Message.Text
			id := update.Message.Chat.ID
			// continue if not mentioned in a group
			if id < 0 && strings.Index(msg, "@"+bot.Self.UserName) < 0 {
				continue
			}
			// clean up the group commands
			if id < 0 {
				msg = strings.Replace(msg, "@"+bot.Self.UserName+" ", "", -1)
				msg = strings.Replace(msg, "@"+bot.Self.UserName, "", -1)
			}

			log.Printf("[%v] %s", id, msg)

			switch msg {
			case "/id":
				{
					SendString(
						id,
						fmt.Sprintf("Your chat id is: %v", id),
						0,
					)
				}
			case "/newmonitor":
				sendWebApp(id)
			default:
				sendWebApp(update.Message.Chat.ID)
			}
		}
	}
}

func sendWebApp(id int64) error {
	msg := tgbotapi.NewMessage(id, "New monitor")
	inlineKb := tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{
		{
			Text: "Add monitor",
			WebApp: &tgbotapi.WebAppInfo{
				URL: "https://tools.lavina.uz:8085/static",
			},
		},
	})
	msg.ReplyMarkup = inlineKb
	_, err := bot.Send(msg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func sendKeyboard(chatId int64) {
	wai := webAppInfo{Url: "https://tools.lavina.uz:8085/static"}
	//wai := webAppInfo{Url: "https://a-webappcontent.stel.com/cafe"}
	button1 := keyboardButton{Text: "My monitors"}
	button2 := keyboardButton{Text: "New monitor", Web_App: &wai}

	keyboard := replyKeyboardMarkUp{Keyboard: [][]keyboardButton{{button1, button2}}, Resize_keyboard: true}

	msg := message{Chat_id: chatId, Text: "Welcome to the bot", Reply_markup: keyboard}
	client := &http.Client{}
	body, _ := json.Marshal(msg)
	req, _ := http.NewRequest("POST", "https://api.telegram.org/bot"+os.Getenv("TG_TOKEN")+"/sendMessage", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	bodybytes, _ := ioutil.ReadAll(res.Body)
	if err != nil || res.StatusCode > 299 {
		fmt.Printf("%s: %s", err, bodybytes)
	}
}

func SendString(id int64, s string, replyId int) {
	msg := tgbotapi.NewMessage(id, escape(s))
	if replyId != 0 {
		msg.ReplyToMessageID = replyId
	}
	msg.ParseMode = "MarkdownV2"
	_, err := bot.Send(msg)
	if err != nil {
		fmt.Println(err)
	}
}

func escape(s string) (r string) {
	var escapeChars = []string{
		"\\", "_", "*", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!",
	}

	// // unEscape the string
	// for _, v := range escapeChars {
	// 	s = strings.Replace(s, "\\"+v, v, -1)
	// }

	// escape the string
	for _, v := range escapeChars {
		s = strings.Replace(s, v, "\\"+v, -1)
	}

	return s
}
