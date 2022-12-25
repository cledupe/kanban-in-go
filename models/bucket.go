package models

type Bucket struct {
	ModelBase
	Cards map[string]bool
}

func NewBucket(name string, color string, userCreator string) *Bucket {
	return &Bucket{
		ModelBase: NewModelBase(name, color, userCreator),
		Cards:     make(map[string]bool),
	}
}

func (b *Bucket) AddCard(cardId string) {
	if exist := b.Cards[cardId]; !exist {
		b.Cards[cardId] = true
		b.SetUpdateAt()
	}
}
