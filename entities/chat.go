package entities

import (
	"time"
)

type Chat struct {
	ID         int       `gorm:"primaryKey"`
	UserID     int       `gorm:"not null"`
	AdminID    int       `gorm:"not null"`
	Message    string    `gorm:"not null;type:text"`
	SenderType string    `gorm:"not null"` // "user" atau "admin"
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type ChatRepositoryInterface interface {
	CreateChat(chat *Chat) error
	GetChatsByUserID(userID int) ([]Chat, error)
	GetChatsByAdminID(adminID int) ([]Chat, error)
	GetChatsBetweenUserAndAdmin(userID, adminID int) ([]Chat, error)
}

type ChatUseCaseInterface interface {
	SendMessage(chat *Chat) error
	GetUserChats(userID int) ([]Chat, error)
	GetAdminChats(adminID int) ([]Chat, error)
	GetConversation(userID, adminID int) ([]Chat, error)
}
