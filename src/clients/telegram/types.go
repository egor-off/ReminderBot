package telegram

type UpdatesResponse struct {
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID		int `json:"update_id"`
	Message	*Message `json:"message,omitempty"`
	CallbackData *CallbackQuery `json:"callback_query,omitempty"`
}

type Message struct {
	Text string `json:"text"`
	From *User `json:"from"`
	Chat Chat `json:"chat"`
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
	CallbackData string `json:"callback_data"`
	// what to send ??
}

type InlineKeyboardMarkup struct {
	Buttons [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type CallbackQuery struct {
	ID string `json:"id"`
	From User `json:"from"`
	Message *Message `json:"message"`
}
