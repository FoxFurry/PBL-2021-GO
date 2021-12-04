package httpresponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func respondWith(c *gin.Context, status int, data interface{}, err interface{}) {
	c.JSON(status, gin.H{
		"data":  data,
		"error": err,
	})
}

func RespondInternalError(c *gin.Context, err interface{}) {
	respondWith(c, http.StatusInternalServerError, nil, err)
}

func RespondNotFound(c *gin.Context, err interface{}) {
	respondWith(c, http.StatusNotFound, nil, err)
}

func RespondUnauthorized(c *gin.Context, err interface{}) {
	respondWith(c, http.StatusUnauthorized, nil, err)
}

func RespondAlreadyExists(c *gin.Context, err interface{}) {
	respondWith(c, http.StatusConflict, nil, err)
}

func RespondOK(c *gin.Context, data interface{}) {
	respondWith(c, http.StatusOK, data, nil)
}
