package repositories

import "cledupe/kanban-in-go/models"

type BucketRepository interface {
	Save(bucket models.Bucket)
	delete(bucket models.Bucket)
	update(bucket models.Bucket)
	findByUser(userId string) []models.Bucket
	findById(Id string) models.Bucket
	findByIdList(IdList []string) []models.Bucket
	addCard(cardId string)
	removeCard(cardId string)
}
