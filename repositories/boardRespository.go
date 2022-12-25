package repositories

import "cledupe/kanban-in-go/models"

type BoardRepository interface {
	save(board models.Bucket)
	delete(board models.Board)
	update(board models.Board)
	findByUser(userId string) []models.Board
	findById(Id string) models.Board
	findByIdList(IdList []string) []models.Board
	addBucket(bucketId string)
	removeBucket(bucketId string)
}
