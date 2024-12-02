package chat

import (
	"e-complaint-api/entities"
	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) entities.ChatRepositoryInterface {
	return &chatRepository{db: db}
}

func (r *chatRepository) CreateChat(chat *entities.Chat) error {
	return r.db.Create(chat).Error
}

func (r *chatRepository) GetChatsByUserID(userID int) ([]entities.Chat, error) {
	var chats []entities.Chat
	err := r.db.Where("user_id = ?", userID).Find(&chats).Error
	return chats, err
}

func (r *chatRepository) GetChatsByAdminID(adminID int) ([]entities.Chat, error) {
	var chats []entities.Chat
	err := r.db.Where("admin_id = ?", adminID).Find(&chats).Error
	return chats, err
}

func (r *chatRepository) GetChatsBetweenUserAndAdmin(userID, adminID int) ([]entities.Chat, error) {
	var chats []entities.Chat
	err := r.db.Where("user_id = ? AND admin_id = ?", userID, adminID).Order("created_at ASC").Find(&chats).Error
	return chats, err
}
