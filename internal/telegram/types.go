package telegram

type (
	webAppInfo struct {
		Url string `json:"url"`
	}

	keyboardButton struct {
		Text    string      `json:"text"`
		Web_App *webAppInfo `json:"web_app,omitempty"`
	}

	replyKeyboardMarkUp struct {
		Keyboard        [][]keyboardButton `json:"keyboard"`
		Resize_keyboard bool               `json:"resize_keyboard"`
	}

	message struct {
		Chat_id      int64               `json:"chat_id"`
		Text         string              `json:"text"`
		Reply_markup replyKeyboardMarkUp `json:"reply_markup"`
	}
)
