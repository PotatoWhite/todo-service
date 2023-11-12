package todo

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	COMPLETE = "complete"
	ACTIVE   = "active"
)

type TodoEntity struct {
	ID        int64     `gorm:"primary_key"`
	Title     string    `gorm:"type:varchar(100)"`
	Status    string    `gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"type:datetime"`
	UpdatedAt time.Time `gorm:"type:datetime"`
}

func (TodoEntity) TableName() string {
	return "todos"
}

type TodoDto struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TodoHandler interface {
	RegisterRoute(r *gin.RouterGroup)
}

type TodoService interface {
	GetAll(status string) ([]TodoEntity, error)
	GetByID(id int64) (*TodoEntity, error)

	Create(todo TodoEntity) (*TodoEntity, error)
	Update(id int64, todo map[string]interface{}) error
	Delete(id int64) error
}
