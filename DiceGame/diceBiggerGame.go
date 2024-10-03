package dicegame

import (
	"fmt"
	"math"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartDiceMoreGame(amount int) (string, *tgbotapi.InlineKeyboardMarkup) {
	InlineKeyBoard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🎲 Да", fmt.Sprintf("DiceBiggerGame/%d", amount)),
		),
	)
	return fmt.Sprintf("Хочешь кинуть кубик и получить %d денег при выигрыше?", amount*2), &InlineKeyBoard
}

func PlayMoreGame(amount int) (string, int) {
	randInt := rand.Intn(6)
	randBotInt := rand.Intn(6)
	randBotInt++
	randInt++
	if randInt > randBotInt {
		return fmt.Sprintf("🤩 Ты выиграл! Полчучено %d денег! (У тебя выпало %d, у бота %d)", amount*2, randInt, randBotInt), amount
	}
	return fmt.Sprintf("😭 Ты проиграл... (У тебя выпало %d, у бота %d)", randInt, randBotInt), int(math.Abs(float64(amount))) * -1
}
