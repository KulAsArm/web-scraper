package services

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramInterface struct {
	Token  string
	ChatID int64
	Bot    *tgbotapi.BotAPI
}

func NewTelegramInterface(token string, chatId int64) *TelegramInterface {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	return &TelegramInterface{
		Bot:    bot,
		Token:  token,
		ChatID: chatId,
	}
}

func (t *TelegramInterface) SendMessage(text string) error {
	msg := tgbotapi.NewMessage(t.ChatID, text)

	_, err := t.Bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
	return err
}
