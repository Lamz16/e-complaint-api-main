package chat

import (
	"e-complaint-api/entities"
)

type chatUseCase struct {
	chatRepo entities.ChatRepositoryInterface
}

func NewChatUseCase(chatRepo entities.ChatRepositoryInterface) entities.ChatUseCaseInterface {
	return &chatUseCase{chatRepo: chatRepo}
}

func (uc *chatUseCase) SendMessage(chat *entities.Chat) error {
	return uc.chatRepo.CreateChat(chat)
}

func (uc *chatUseCase) GetUserChats(userID int) ([]entities.Chat, error) {
	return uc.chatRepo.GetChatsByUserID(userID)
}

func (uc *chatUseCase) GetAdminChats(adminID int) ([]entities.Chat, error) {
	return uc.chatRepo.GetChatsByAdminID(adminID)
}

func (uc *chatUseCase) GetConversation(userID, adminID int) ([]entities.Chat, error) {
	return uc.chatRepo.GetChatsBetweenUserAndAdmin(userID, adminID)
}
