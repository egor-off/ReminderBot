package telegram

import "src/clients/telegram"

var (
	//Default sending
	defaultKeyboard = telegram.InlineKeyboardMarkup{
		Buttons: [][]telegram.InlineKeyboardButton{
			{
				helpButton,
				randomButton,
				newRemindButton,
			},
		},
	}

	//Sending with rnd answer
	AfterRndKeyboard = telegram.InlineKeyboardMarkup{
		Buttons: [][]telegram.InlineKeyboardButton{
			{
				saveURLButton,
				deleteURLButton,
			},
		},
	}

)
