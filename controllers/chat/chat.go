package chat

import (
	"e-complaint-api/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type ChatController struct {
	chatUsecase entities.ChatUseCaseInterface
}

func NewChatController(chatUsecase entities.ChatUseCaseInterface) *ChatController {
	return &ChatController{chatUsecase: chatUsecase}
}

func (c *ChatController) SendMessage(ctx echo.Context) error {
	var request struct {
		UserID     int    `json:"userID"`
		AdminID    int    `json:"adminID"`
		Message    string `json:"message"`
		SenderType string `json:"senderType"` // "user" atau "admin"
	}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if request.UserID == 0 || request.AdminID == 0 || request.Message == "" || request.SenderType == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
	}

	chat := entities.Chat{
		UserID:     request.UserID,
		AdminID:    request.AdminID,
		Message:    request.Message,
		SenderType: request.SenderType,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := c.chatUsecase.SendMessage(&chat)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send message"})
	}

	return ctx.JSON(http.StatusOK, chat)
}

func (c *ChatController) GetConversation(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.QueryParam("user_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	adminID, err := strconv.Atoi(ctx.QueryParam("admin_id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid admin ID"})
	}

	chats, err := c.chatUsecase.GetConversation(userID, adminID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch chats"})
	}

	return ctx.JSON(http.StatusOK, chats)
}
