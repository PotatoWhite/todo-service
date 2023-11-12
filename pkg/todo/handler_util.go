package todo

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func entityToDto(todo TodoEntity) TodoDto {
	return TodoDto{
		ID:        todo.ID,
		Title:     todo.Title,
		Status:    todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func dtoToEntity(todo TodoDto) TodoEntity {
	return TodoEntity{
		ID:        todo.ID,
		Title:     todo.Title,
		Status:    todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func getID(c *gin.Context) int64 {
	// id from path
	if id, err := strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		panic(ErrInvalidID)
	} else {
		return id
	}
}
