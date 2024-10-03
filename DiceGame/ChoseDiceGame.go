package dicegame

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ChoseDiceGame(amount int) (string, *tgbotapi.InlineKeyboardMarkup) {
	InlineKeyBoard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🎲 У кого больше", fmt.Sprintf("BiggerDiceGameStart/%d", amount)),
			tgbotapi.NewInlineKeyboardButtonData("🎲 Угадай", fmt.Sprintf("GuessDiceGameStart/%d", amount)),
		),
	)
	return fmt.Sprintf("Выбери игру на которую хочешь играть со ставкой %d", amount), &InlineKeyBoard
}
