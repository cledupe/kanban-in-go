package models

import (
	"time"

	"github.com/google/uuid"
)

type ModelBase struct {
	ID          string
	Name        string
	Color       string
	UserCreator string
	CreatedAt   time.Time
	UpdateAt    time.Time
}

func NewModelBase(name string, color string, userCreator string) ModelBase {
	id := uuid.New().String()
	now := time.Now()
	return ModelBase{
		ID:          id,
		Name:        name,
		Color:       color,
		UserCreator: userCreator,
		CreatedAt:   now,
		UpdateAt:    now,
	}
}

func (m *ModelBase) SetName(name string) {
	m.Name = name
	m.SetUpdateAt()
}

func (m *ModelBase) SetColor(color string) {
	m.Color = color
	m.SetUpdateAt()
}

func (m *ModelBase) SetUpdateAt() {
	m.UpdateAt = time.Now()
}
