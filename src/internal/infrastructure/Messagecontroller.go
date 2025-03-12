package infrastructure

import (
	"encoding/json"
	"net/http"
	"pedidos/src/internal/application"
)

type MessageController struct {
	getMessageUseCase    *application.GetMessageUseCase
	createMessageUseCase *application.CreateMessageUseCase
}

func NewMessageController(
	getMessageUseCase *application.GetMessageUseCase,
	createMessageUseCase *application.CreateMessageUseCase,
) *MessageController {
	return &MessageController{
		getMessageUseCase:    getMessageUseCase,
		createMessageUseCase: createMessageUseCase,
	}
}

func (c *MessageController) GetMessages(w http.ResponseWriter, r *http.Request) {
	message, err := c.getMessageUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
