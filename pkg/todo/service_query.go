package todo

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (s todoService) GetAll(status string) ([]TodoEntity, error) {
	var todos []TodoEntity

	if status == "" {
		if err := s.db.Find(&todos).Error; err != nil {
			return nil, err
		}
		return todos, nil
	} else {
		if err := s.db.Where(fmt.Sprintf("status = %s", status)).Find(&todos).Error; err != nil {
			return nil, err
		}
		return todos, nil
	}
}

func (s todoService) GetByID(id int64) (*TodoEntity, error) {
	var todo *TodoEntity

	if tx := s.db.Where("id = ?", id).First(&todo); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return todo, nil
}
