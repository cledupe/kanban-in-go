package models

type Card struct {
	ModelBase
	Content      string
	Priority     int64
	AssignedUser string
}

func NewCard(name string, content string, priority int64, userCreator string, assignedUser string, color string) Card {

	return Card{
		ModelBase:    NewModelBase(name, color, userCreator),
		Content:      content,
		Priority:     priority,
		AssignedUser: assignedUser,
	}
}

func (card *Card) SetContent(content string) {
	card.Content = content
	card.SetUpdateAt()
}

func (card *Card) SetPriority(priority int64) {
	card.Priority = priority
	card.SetUpdateAt()
}

func (card *Card) SetAssignedUser(assignedUser string) {
	card.AssignedUser = assignedUser
	card.SetUpdateAt()
}
