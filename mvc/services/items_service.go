package services

import (
	"net/http"

	"github.com/Go-Lang/mvc/domain"
	"github.com/Go-Lang/mvc/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
