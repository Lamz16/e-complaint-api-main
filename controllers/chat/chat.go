package chat

import (
	"e-complaint-api/entities"
	"github.com/labstack/echo/v4"
	"log"
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
		log.Println("Invalid input data:", err)
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
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
	}

	err := c.chatUsecase.SendMessage(&chat)
	if err != nil {
		log.Println("Failed to send message:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send message"})
	}

	log.Printf("Message sent successfully: %v", chat)
	return ctx.JSON(http.StatusOK, chat)
}

func (c *ChatController) GetConversation(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.QueryParam("user_id"))
	if err != nil || userID <= 0 {
		log.Println("Invalid user ID:", ctx.QueryParam("user_id"))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	adminID, err := strconv.Atoi(ctx.QueryParam("admin_id"))
	if err != nil || adminID <= 0 {
		log.Println("Invalid admin ID:", ctx.QueryParam("admin_id"))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid admin ID"})
	}

	chats, err := c.chatUsecase.GetConversation(userID, adminID)
	if err != nil {
		log.Println("Failed to fetch chats:", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch chats"})
	}

	log.Printf("Fetched conversation for userID %d and adminID %d", userID, adminID)
	return ctx.JSON(http.StatusOK, chats)
}
