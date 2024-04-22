package telegram

import "strings"

var (
	msgHelp = `I can save links! 🔗

For saving links just send me link in a message. ✍

Press the "` + randomText + `" button and you'll get random link from messages you have sent before.`

	msgHello = "Hello! This is simple autoreminder and linksaver bot!\n\n" + msgHelp

	msgDeafult = "Push the button to get the result ⏬"

	msgUnknownCommand = "Unknown command 🫠\n\nPlease, send me a link to save.\nOr " + strings.ToLower(msgDeafult)
	msgNoSavedURL = "No saved URL 😬\n\n" + msgDeafult
	msgAllreadyExists = "This URL already exists 🤌🏽\n\n" + msgDeafult
	msgSaved = "Saved 🤞🏻\n\n" + msgDeafult
	msgDeleted = "Link was deleted 🚮\n\n" + msgDeafult
)
