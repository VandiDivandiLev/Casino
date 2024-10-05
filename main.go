package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	centgame "telegram-balance-bot/CentGame"
	dicegame "telegram-balance-bot/DiceGame"
	reply "telegram-balance-bot/Reply"
	slotmachinegame "telegram-balance-bot/SlotMachineGame"
	"telegram-balance-bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type User struct {
	ID      int64 `json:"id"`
	Balance int   `json:"balance"`
}

type UserMap map[int64]User

func main() {
	bot, err := tgbotapi.NewBotAPI(config.TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Load users from JSON file or create an empty map
	var users UserMap
	if err := loadUsers("users.json", &users); err != nil {
		// If file not found, create an empty map
		users = make(UserMap)
		users[int64(bot.Self.ID)] = User{ID: int64(bot.Self.ID), Balance: 0}
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {
			// Construct a new message from the given chat ID and containing
			// the text that we received.
			switch update.Message.Command() {
			case "start":
				reply.Repl(update, "👋 Привет! Я бот для игры в казино", nil, bot)
			case "help":
				reply.Repl(update, "👉 Мои игры: \n<code>монетка</code>\n<code>кубик</code>", nil, bot)
			default:
				splitText := strings.Split(strings.ToLower(update.Message.Text), " ")
				switch splitText[0] {
				case "баланс":
					user := update.Message.From.ID
					balance, ok := users[user]

					if !ok {
						// User doesn't exist, create a new entry with balance 0
						balance = User{ID: user, Balance: 0}
						users[user] = balance
					}

					fmt.Println(ok, users) // Print debugging information

					// Format the balance string for the message
					var balanceString string = fmt.Sprintf("Ваш баланс: %d", balance.Balance)
					fmt.Println(balanceString)
					reply.Repl(update, balanceString, nil, bot)
				case "монетка":
					amount := 1
					if len(splitText) >= 2 {
						amount, _ = strconv.Atoi(splitText[1])
					}
					if amount <= 0 {
						amount = 1
					}
					replString, Keybord := centgame.StartCentGame(amount)
					reply.Repl(update, replString, Keybord, bot)
				case "автомат":
					amount := 1
					if len(splitText) >= 2 {
						amount, _ = strconv.Atoi(splitText[1])
					}
					if amount <= 0 {
						amount = 1
					}
					replString, Keybord := slotmachinegame.StartSlotMachineGame(amount)
					reply.Repl(update, replString, Keybord, bot)
				case "кубик":
					amount := 1
					if len(splitText) >= 2 {
						amount, _ = strconv.Atoi(splitText[1])
					}
					if amount <= 0 {
						amount = 1
					}
					replString, Keybord := dicegame.ChoseDiceGame(amount)
					reply.Repl(update, replString, Keybord, bot)
				}
			}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.

			// And finally, send a message containing the data received.
			user := update.CallbackQuery.From.ID
			data := strings.Split(update.CallbackQuery.Data, "/")
			Info, ok := users[user]
			if !ok {
				users[user] = User{ID: user, Balance: 0}
			}
			switch data[0] {
			case "BiggerDiceGameStart":
				cost, _ := strconv.Atoi(data[1])
				text, keys := dicegame.StartDiceMoreGame(cost)
				reply.Repl(update, text, keys, bot)
			case "GuessDiceGameStart":
				cost, _ := strconv.Atoi(data[1])
				text, keys := dicegame.StartDiceGuessGame(cost)
				reply.Repl(update, text, keys, bot)
			case "CentGame":
				cost, _ := strconv.Atoi(data[2])
				chose, _ := strconv.ParseBool(data[1])
				if Info.Balance >= cost {
					text, amount := centgame.PlayGame(cost, chose)

					NewBal := Info.Balance + amount
					users[user] = User{ID: user, Balance: NewBal}
					reply.Repl(update, text, nil, bot)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, text)
					if _, err := bot.Request(callback); err != nil {
						panic(err)
					}
				} else {
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "❌ Недостаточно средств")
					if _, err := bot.Request(callback); err != nil {
						panic(err)
					}
				}
			case "DiceGuessGame":
				cost, _ := strconv.Atoi(data[2])
				chose, _ := strconv.Atoi(data[1])
				if Info.Balance >= cost {
					text, amount := dicegame.PlayGuessGame(cost, chose)

					NewBal := Info.Balance + amount
					users[user] = User{ID: user, Balance: NewBal}
					reply.Repl(update, text, nil, bot)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, text)
					if _, err := bot.Request(callback); err != nil {
						panic(err)
					}
				} else {
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "❌ Недостаточно средств")
					if _, err := bot.Request(callback); err != nil {
						panic(err)
					}
				}
			case "DiceBiggerGame":
				cost, _ := strconv.Atoi(data[1])
				if Info.Balance >= cost {
					text, amount := dicegame.PlayMoreGame(cost)

					NewBal := Info.Balance + amount
					users[user] = User{ID: user, Balance: NewBal}
					reply.Repl(update, text, nil, bot)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, text)
					if _, err := bot.Request(callback); err != nil {
						panic(err)
					}
				} else {
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "❌ Недостаточно средств")
					if _, err := bot.Request(callback); err != nil {
						panic(err)
					}
				}
			case "SlotMachineGame":
				cost, _ := strconv.Atoi(data[1])
				if Info.Balance >= cost {
					text, amount := slotmachinegame.PlaySlotMachineGame(cost)

					NewBal := Info.Balance + amount
					users[user] = User{ID: user, Balance: NewBal}
					reply.Repl(update, text, nil, bot)
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, text)
					if _, err := bot.Request(callback); err != nil {
						panic(err)
					}
				} else {
					callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "❌ Недостаточно средств")
					if _, err := bot.Request(callback); err != nil {
						panic(err)
					}
				}
			}

		}

		// Save users to JSON file after each update
		if err := saveUsers("users.json", users); err != nil {
			log.Fatal(err)
		}
	}
}

// Load users from JSON file
func loadUsers(filename string, users *UserMap) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, users)
}

// Save users to JSON file
func saveUsers(filename string, users UserMap) error {
	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}
