package dicegame

import (
	"fmt"
	"math"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartDiceGuessGame(amount int) (string, *tgbotapi.InlineKeyboardMarkup) {
	InlineKeyBoard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🎲 1", fmt.Sprintf("DiceGuessGame/1/%d", amount)),
			tgbotapi.NewInlineKeyboardButtonData("🎲 2", fmt.Sprintf("DiceGuessGame/2/%d", amount)),
			tgbotapi.NewInlineKeyboardButtonData("🎲 3", fmt.Sprintf("DiceGuessGame/3/%d", amount)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🎲 4", fmt.Sprintf("DiceGuessGame/4/%d", amount)),
			tgbotapi.NewInlineKeyboardButtonData("🎲 5", fmt.Sprintf("DiceGuessGame/5/%d", amount)),
			tgbotapi.NewInlineKeyboardButtonData("🎲 6", fmt.Sprintf("DiceGuessGame/6/%d", amount)),
		),
	)
	return fmt.Sprintf("Хочешь кинуть кубик и получить %d денег при выигрыше. Если да, то выбери:", amount*6), &InlineKeyBoard
}

func PlayGuessGame(amount int, Part int) (string, int) {
	randInt := rand.Intn(6)
	randInt++
	if randInt == Part {
		return fmt.Sprintf("🤩 Ты выиграл! Полчучено %d денег! (Выпало %d)", amount*6, randInt), amount * 5
	}
	return fmt.Sprintf("😭 Ты проиграл... (Выпало %d, ставка на %d)", randInt, Part), int(math.Abs(float64(amount))) * -1
}
