package todo

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func errorHandler(err error, c *gin.Context) {
	if errors.Is(err, sql.ErrNoRows) {
		c.AbortWithError(http.StatusNoContent, err)
		return
	} else if errors.Is(err, ErrInvalidID) {
		c.AbortWithError(http.StatusNoContent, err)
		return
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func recoverAndErrorResponse(c *gin.Context) {
	func() {
		if r := recover(); r != nil {
			err := errors.Unwrap(r.(error))
			errorHandler(err, c)
		}
	}()
}
