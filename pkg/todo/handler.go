package todo

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidID = errors.New("invalid id")
)

type todoHandler struct {
	service TodoService
}

func NewHandler(service TodoService) TodoHandler {
	return &todoHandler{service: service}
}

func (t todoHandler) RegisterRoute(r *gin.RouterGroup) {
	r.GET("/", t.getAll)
	r.GET("/:id", t.getByID)
	r.POST("/", t.create)
	r.PATCH("/:id", t.update)
	r.DELETE("/:id", t.delete)
}
