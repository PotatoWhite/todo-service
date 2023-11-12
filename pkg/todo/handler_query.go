package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t todoHandler) getAll(c *gin.Context) {
	defer recoverAndErrorResponse(c)

	// status from query
	status := c.Query("status")

	all, err := t.service.GetAll(status)
	if err != nil {
		panic(err)
	}

	var result []TodoDto
	for _, todo := range all {
		result = append(result, entityToDto(todo))
	}

	c.JSON(http.StatusOK, result)
}

func (t todoHandler) getByID(c *gin.Context) {
	defer recoverAndErrorResponse(c)

	id := getID(c)

	a, err := t.service.GetByID(id)
	if err != nil {
		panic(err)
	} else if a == nil {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, entityToDto(*a))
}
