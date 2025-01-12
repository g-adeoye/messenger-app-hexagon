package ports

import "Messenger/internal/core/domain"

type MessengerService interface {
	SaveMessage(message domain.Message) error
	ReadMessage(id string) (*domain.Message, error)
	ReadMessages() ([]*domain.Message, error)
}

type MessageRepository interface {
	SaveMessage(message domain.Message) error
	ReadMessage(id string) (*domain.Message, error)
	ReadMessages() ([]*domain.Message, error)
}
