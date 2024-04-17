package telegram

import "src/clients/telegram"

var (
	// Default sending
	defaultKeyboard = &telegram.InlineKeyboardMarkup{
		Buttons: [][]telegram.InlineKeyboardButton{
			{
				helpButton,
				randomButton,
				// newRemindButton,
			},
		},
	}

	// Sending with rnd answer
	afterRndKeyboard = &telegram.InlineKeyboardMarkup{
		Buttons: [][]telegram.InlineKeyboardButton{
			{
				deleteURLButton,
				saveURLButton,
				nextURLButton,
			},
		},
	}

	// Remind buttons
	remindKeyboard = &telegram.InlineKeyboardMarkup{
		Buttons: [][]telegram.InlineKeyboardButton{
			{
				textRemindButton,
				dateRemindButton,
				periodRemindButton,
				cancelRemindButton,
			},
		},
	}
)
