package whatsappbot

import (
	"database/sql"
	"time"

	"go.mau.fi/whatsmeow/types"
)

type ConversationContext struct {
	UserInfo     UserInfo
	UserExisted  bool
	CurrentOrder CustomerOrder
	MessageBody  string
	SenderJID    types.JID
	DBReadTime   time.Time
}

func NewConversationContext(db *sql.DB, senderNumber, messagebody string, isAutoInc bool) *ConversationContext {
	userInfo, curOrder, userExisted := NewUserInfo(db, senderNumber, isAutoInc)
	context := &ConversationContext{
		UserInfo:     userInfo,
		UserExisted:  userExisted,
		CurrentOrder: curOrder,
		MessageBody:  messagebody,
		SenderJID:    types.NewJID(userInfo.CellNumber, "s.whatsapp.net"),
		DBReadTime:   time.Now(),
	}

	return context
}
