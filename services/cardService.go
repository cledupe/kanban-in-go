package services

import (
	"cledupe/kanban-in-go/models"
	"cledupe/kanban-in-go/repositories"
)

type CardService struct {
	repo repositories.CardRepository
}

func (c *CardService) Create(name string, content string, priority int64, userCreator string, assignedUser string, color string) {

	card := models.NewCard(name, content, priority, userCreator, assignedUser, color)
	c.repo.Save(card)
}
