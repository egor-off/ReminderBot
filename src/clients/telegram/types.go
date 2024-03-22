package telegram

type UpdatesResponse struct {
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID		int `json:"update_id"`
	Message	*IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From User `json:"from"`
	Chat Chat `json:"chat"`
}

type SendingMessage struct {
	ChatID int `json:"chat_id"`
	Text string `json:"text"`
	Buttons *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type User struct {
	UserName string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}

type InlineKeyboardButton struct {
	Text string `json:"text"`
	// what to send ??
}

type InlineKeyboardMarkup struct {
	Buttons [][]InlineKeyboardButton `json:"inline_keyboard"`
}
