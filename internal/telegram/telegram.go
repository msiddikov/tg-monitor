package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func StartTelegram() {
	godotenv.Load(".env")
	token := os.Getenv("TG_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			//bot.Send(msg)
			sendKeyboard(update.Message.Chat.ID)
		}
	}
}

func sendKeyboard(chatId int64) {
	wai := webAppInfo{Url: "https://tools.lavina.uz:8085/static"}
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
