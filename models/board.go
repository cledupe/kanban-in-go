package models

type Board struct {
	ModelBase
	Buckets map[string]bool
}

func NewBoard(name string, color string, userCreator string) *Board {
	return &Board{
		ModelBase: NewModelBase(name, color, userCreator),
		Buckets:   make(map[string]bool),
	}
}

func (b *Board) AddBucket(bucketId string) {
	if exist := b.Buckets[bucketId]; !exist {
		b.Buckets[bucketId] = true
		b.SetUpdateAt()
	}
}
