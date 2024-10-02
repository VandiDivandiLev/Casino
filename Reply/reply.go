package reply

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Repl sends a message to the chat with an optional keyboard, splitting the text if it's too long.
func Repl(update tgbotapi.Update, text string, keyboard *tgbotapi.InlineKeyboardMarkup, Bot *tgbotapi.BotAPI) error {
	// Check for valid text input
	if text == "" {
		return fmt.Errorf("text cannot be empty")
	}

	// Split the message into parts if it exceeds 4096 characters
	messageParts := strings.Split(text, "\n")

	// Determine the ID of the message to reply to
	replyToMessageId := 0
	if update.Message != nil && update.Message.MessageID != 0 {
		replyToMessageId = update.Message.MessageID
	} else if update.Message != nil && update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.MessageID != 0 {
		replyToMessageId = update.Message.ReplyToMessage.MessageID
	}

	// Send the message parts iteratively
	for i := 0; i < len(messageParts); i++ {
		part := messageParts[i]
		// Build the message options
		message := update.Message
		if message == nil {
			message = update.CallbackQuery.Message
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, part)
		msg.ParseMode = "HTML"
		msg.ReplyToMessageID = replyToMessageId
		msg.DisableWebPagePreview = true

		// Add the keyboard if it's provided and it's the last part
		if keyboard != nil && i == len(messageParts)-1 {
			msg.ReplyMarkup = keyboard
		}

		// Send the message part
		Bot.Send(msg)
	}

	return nil
}
