package lib

type Message struct {
	Date               int64               `json:"date"`
	Chat               *Chat               `json:"chat"`
	MessageId          int64               `json:"message_id"`
	From               *From               `json:"from"`
	Text               string              `json:"text"`
	ReplyMarkup        *ReplyMarkup        `json:"reply_markup"`
	ReplyToMessage     *Message            `json:"reply_to_message"`
	NewChatParticipant *NewChatParticipant `json:"new_chat_participant"`
	NewChatMembers     *[]NewChatMembers   `json:"new_chat_members"`
	SenderChat         *SenderChat         `json:"sender_chat"`
	NewChatMember      *NewChatMember      `json:"new_chat_member"`
}
