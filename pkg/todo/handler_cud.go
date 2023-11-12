package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t todoHandler) create(c *gin.Context) {
	defer recoverAndErrorResponse(c)

	// todo from body
	var todo TodoDto
	if err := c.ShouldBindJSON(&todo); err != nil {
		panic(err)
	}

	saved, err := t.service.Create(dtoToEntity(todo))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, entityToDto(*saved))
}

func (t todoHandler) update(c *gin.Context) {
	defer recoverAndErrorResponse(c)

	var fields map[string]interface{}
	if err := c.ShouldBindJSON(&fields); err != nil {
		panic(err)
	}

	id := getID(c)

	if fields["status"] == nil {
		panic(ErrInvalidID)
	} else if fields["status"] == COMPLETE || fields["status"] == ACTIVE {
		panic(ErrInvalidID)
	}

	if err := t.service.Update(id, fields); err != nil {
		t.service.Update(id, fields)
	}

	c.JSON(http.StatusOK, nil)
}

func (t todoHandler) delete(c *gin.Context) {
	defer recoverAndErrorResponse(c)

	id := getID(c)
	if err := t.service.Delete(id); err != nil {
		panic(err)
	}

	c.JSON(http.StatusNoContent, nil)
}
