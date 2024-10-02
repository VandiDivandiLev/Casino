package centgame

import (
	"fmt"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCentGame(amount int) (string, *tgbotapi.InlineKeyboardMarkup) {
	InlineKeyBoard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🦅 Орел", fmt.Sprintf("CentGame/true/%d", amount)),
			tgbotapi.NewInlineKeyboardButtonData("🪙 Решка", fmt.Sprintf("CentGame/false/%d", amount)),
		),
	)
	return fmt.Sprintf("Хочешь подбросить монетку и получить %d денег при выигрыше. Если да, то выбери:", amount*2), &InlineKeyBoard
}

func PlayGame(amount int, Part bool) (string, int) {
	randInt := rand.Intn(2)
	randomBool := false
	if randInt == 1 {
		randomBool = true
	}
	if randomBool == Part {
		return fmt.Sprintf("🤩 Ты выиграл! Полчучено %d денег", amount*2), amount
	}
	return "😭 Ты проиграл...", 0
}
