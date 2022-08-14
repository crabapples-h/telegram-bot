package eneity

type Bot struct {
	Id          string `json:"id" orm:"column(id);pk"`
	ChatId      int64  `json:"ChatId" orm:"column(chat_id)"`
	GroupName   string `json:"groupName" orm:"column(group_name)"`
	BotName     string `json:"botName" orm:"column(bot_name)"`
	Token       string `json:"token" orm:"column(token)"`
	Username    string `json:"username" orm:"column(username);unique;index;description(这是状态字段)"`
	CallbackUrl string `json:"callbackUrl" orm:"column(callback_url)"`
	DelFlag     int8   `json:"delFlag" orm:"column(del_flag)"`
}
