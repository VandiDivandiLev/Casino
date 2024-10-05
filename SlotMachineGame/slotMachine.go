package slotmachinegame

import (
	"fmt"
	"math"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartSlotMachineGame(amount int) (string, *tgbotapi.InlineKeyboardMarkup) {
	InlineKeyBoard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🎰 Да", fmt.Sprintf("SlotMachineGame/%d", amount)),
		),
	)
	return fmt.Sprintf("Хочешь крутануть автомат и получить %d от до %d денег при выигрыше?", amount*10, amount*100), &InlineKeyBoard
}

func PlaySlotMachineGame(amount int) (string, int) {
	emojis := []string{"🍒", "🍋", "🍊", "🍉", "🔔", "⭐", "7️⃣", "💰", "🍀"}
	index1 := rand.Intn(len(emojis))
	index2 := rand.Intn(len(emojis))
	index3 := rand.Intn(len(emojis))
	slot1 := emojis[index1]
	slot2 := emojis[index2]
	slot3 := emojis[index3]
	multiple := 100
	if slot1 == "7️⃣" {
		multiple = 100
	} else if slot1 == "🍀" || slot1 == "💰" || slot1 == "⭐" || slot1 == "🔔" {
		multiple = 50
	}
	slotrel := fmt.Sprintf("%s %s %s", slot1, slot2, slot3)
	if slot1 == slot2 && slot2 == slot3 {
		return fmt.Sprintf("🤩 Ты выиграл! Получено %d денег (Выпало %s)", amount*multiple, slotrel), amount * (multiple - 1)
	} else if slot1 == slot2 || slot1 == slot3 || slot2 == slot3 {
		return fmt.Sprintf("🤨 Выпало недостаточно для выигрша! (Выапало %s)", slotrel), 0
	} else {
		return fmt.Sprintf("😭 Ты проиграл... (Выпало %s)", slotrel), int(math.Abs(float64(amount))) * -1

	}
}
