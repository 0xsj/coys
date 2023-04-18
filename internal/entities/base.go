package entities

import (
	"time"

	"github.com/google/uuid"
)

type Interface interface {
	GenerateID()
	SetCreatedAt()
	SetUpdatedAt()
	TableName() string
	GetMap() map[string]interface{}
	GetFilterID() map[string]interface{}
}

type Base struct {
	ID        uuid.UUID `json:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GenerateID(b *Base) {
	b.ID = uuid.New()
}

func SetCreatedAt(b *Base) {
	b.CreatedAt = time.Now()
}

func SetUpdatedAt(b *Base) {}

func GetTimeFormat(b *Base) string {
	return "2006-01-02T15:04:05-0700"
}
