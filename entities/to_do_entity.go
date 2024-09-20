package entities

import (
	"fmt"
	"gorm.io/gorm"
)

type TODO struct {
	Message     string `json:"message"`
	IsImportant bool   `json:"is_important" gorm:"default:false;type:bool"`

	gorm.Model
}

func NewTODO(message string, isImportant bool) *TODO {
	return &TODO{
		Message:     message,
		IsImportant: isImportant,
	}
}

func (t *TODO) String() string {
	if t.IsImportant {
		return fmt.Sprintf("%v\n★ \"%v\"\n%v", t.ID, t.Message, t.UpdatedAt)
	}

	return fmt.Sprintf("%v\n☆ \"%v\"\n%v", t.ID, t.Message, t.UpdatedAt)
}
