package eneity

type Command struct {
	Id           string `json:"id" orm:"column(id);pk"`
	Command      string `json:"command" orm:"column(command)"`
	Description  string `json:"description" orm:"column(description)"`
	ReplyMessage string `json:"replyMessage" orm:"column(reply_message)"`
	DelFlag      int8   `json:"delFlag" orm:"column(del_flag)"`
}
