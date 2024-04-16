package telegram

//UpdateResponse is the answer for getUpdates response, representing ok bool and slice of updates.
type UpdatesResponse struct {
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}

//Update is the one of Updates from slice, we got from response getU[dates.
type Update struct {
	ID		int `json:"update_id"`
	Message	*Message `json:"message,omitempty"`
	CallbackData *CallbackQuery `json:"callback_query,omitempty"`
}

//Message is type representig the Message we can get in Update.
type Message struct {
	Text string `json:"text"`
	From *User `json:"from"`
	Chat Chat `json:"chat"`
	Buttons *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

//User is type representing the User from we get Updates.
type User struct {
	UserName string `json:"username"`
}

//Chat is type representing the Chat where we can get the dialog ID.
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
	Data string `json:"data"`
}
