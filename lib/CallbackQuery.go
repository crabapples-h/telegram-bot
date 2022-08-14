package lib

type CallbackQuery struct {
	Id           string  `json:"id"`
	Data         string  `json:"data"`
	From         From    `json:"from"`
	Message      Message `json:"message"`
	ChatInstance string  `json:"chat_instance"`
}
