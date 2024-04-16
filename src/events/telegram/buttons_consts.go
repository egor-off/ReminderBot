package telegram

import "src/clients/telegram"

var (
	// Default buttons.
	// Help.
	helpText = "Help ⁉️"
	helpData = "/help"
	helpButton = telegram.InlineKeyboardButton{
		Text: helpText,
		CallbackData: helpData,
	}

	// Random.
	rndText = "Random 🎲"
	rndData = "/rnd"
	randomButton = telegram.InlineKeyboardButton{
		Text: rndText,
		CallbackData: rndData,
	}

	// New remind.
	newRemindText = "New remind 📝"
	newRemindData = "/newRemind"
	newRemindButton = telegram.InlineKeyboardButton{
		Text: newRemindText,
		CallbackData: newRemindData,
	}

	// URL buttons.
	// TODO: mb add some more? like helpURL and return?
	// Delete URL.
	deleteUrlText = "Delete 🗑"
	deleteUrlData = "/deleteURL"
	deleteURLButton = telegram.InlineKeyboardButton{
		Text: deleteUrlText,
		CallbackData: deleteUrlData,
	}

	// Keep save URL.
	keepSaveText = "Keep save 🫡"
	keepSaveData = "/saveURL"
	saveURLButton = telegram.InlineKeyboardButton{
		Text: keepSaveText,
		CallbackData: keepSaveData,
	}

	// Reminds.
	// Text remind.
	textRemindText = "Add text 💬"
	textRemindData = "/textRemind"
	textRemindButton = telegram.InlineKeyboardButton{
		Text: textRemindText,
		CallbackData: textRemindData,
	}

	// Remind date
	dateRemindText = "Add date 📅"
	dateRemindData = "/dateRemind"
	dateRemindButton = telegram.InlineKeyboardButton{
		Text: dateRemindText,
		CallbackData: dateRemindData,
	}

	// Remind period
	periodRemindText = "Add period 🕑"
	periodRemindData = "/periodRemind"
	periodRemindButton = telegram.InlineKeyboardButton{
		Text: periodRemindText,
		CallbackData: periodRemindData,
	}

	// Remind cancel
	cancelRemindText = "Cancel ⛔️"
	cancelRemindData = "/cancelRemind"
	cancelRemindButton = telegram.InlineKeyboardButton{
		Text: cancelRemindText,
		CallbackData: cancelRemindData,
	}
)
