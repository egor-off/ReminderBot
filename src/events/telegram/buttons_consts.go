package telegram

import "src/clients/telegram"

var (
	// Help
	helpText = "Help ⁉️"
	helpData = "/help"
	helpButton = telegram.InlineKeyboardButton{
		Text: helpText,
		CallbackData: helpData,
	}

	// Random
	rndText = "Random 🎲"
	rndData = "/rnd"
	randomButton = telegram.InlineKeyboardButton{
		Text: rndText,
		CallbackData: rndData,
	}

	// New remind
	newRemindText = "New remind 📝"
	newRemindData = "/new_remind"
	newRemindButton = telegram.InlineKeyboardButton{
		Text: newRemindText,
		CallbackData: newRemindData,
	}


	// Delete URL
	deleteUrlText = "Delete 🗑"
	deleteUrlData = "/deleteURL"
	deleteURLButton = telegram.InlineKeyboardButton{
		Text: deleteUrlText,
		CallbackData: deleteUrlData,
	}

	// Keep save
	keepSaveText = "Keep save 🫡"
	keepSaveData = "/saveURL"
	saveURLButton = telegram.InlineKeyboardButton{
		Text: keepSaveText,
		CallbackData: keepSaveData,
	}
)
