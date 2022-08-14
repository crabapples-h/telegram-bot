package form

import "telegram-bot/lib"

type CallBackData struct {
	UpdateId      int64             `json:"update_id"`
	CallbackQuery lib.CallbackQuery `json:"callback_query"`
	Message       lib.Message       `json:"message"`
}
