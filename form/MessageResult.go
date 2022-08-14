package form

import "telegram-bot/lib"

type MessageResult struct {
	Ok     bool        `json:"ok"`
	Result lib.Message `json:"result"`
}
