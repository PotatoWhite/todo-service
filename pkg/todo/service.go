package todo

import (
	"gorm.io/gorm"
)

type todoService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) TodoService {
	return &todoService{db: db}
}
