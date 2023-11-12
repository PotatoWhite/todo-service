package todo

import (
	"errors"
	"gorm.io/gorm"
)

func (s todoService) Create(todo TodoEntity) (*TodoEntity, error) {
	if todo.Status == "" {
		todo.Status = ACTIVE
	}

	if err := s.db.Create(&todo).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (s todoService) Update(id int64, todo map[string]interface{}) error {
	var todoEntity TodoEntity

	if tx := s.db.Where("id = ?", id).First(&todoEntity); tx.Error != nil {
		return tx.Error
	}

	if tx := s.db.Model(&todoEntity).Updates(todo); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (s todoService) Delete(id int64) error {
	var todo TodoEntity

	if tx := s.db.Where("id = ?", id).Delete(&todo); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return tx.Error
	}
	return nil
}
