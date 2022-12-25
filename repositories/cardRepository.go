package repositories

import "cledupe/kanban-in-go/models"

type CardRepository interface {
	Save(card models.Card)
	delete(card models.Card)
	update(card models.Card)
	findByUser(userId string) []models.Card
	findById(id string) models.Card
	findByIdList(idList []string)
}
