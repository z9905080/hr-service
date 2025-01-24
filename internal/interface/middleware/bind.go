package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type C struct {
	c *gin.Context
}

func HandleFunc(handler func(c *C)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&C{c: c})
	}
}

func (c *C) GetCtx() *gin.Context {
	return c.c
}

func (c *C) GetParam(key string) string {
	return c.c.Param(key)
}

func (c *C) GetQuery(key string) string {
	return c.c.Query(key)
}

func (c *C) BindJSON(obj interface{}) error {

	// bind uri
	if err := c.c.ShouldBindUri(obj); err != nil {
		return err
	}

	// bind json
	if err := c.c.ShouldBindJSON(obj); err != nil {
		return err
	}

	return nil
}

func (c *C) BindQuery(obj interface{}) error {

	// bind query
	if err := c.c.ShouldBindQuery(obj); err != nil {
		return err
	}

	return nil
}

func (c *C) BindUri(obj interface{}) error {

	// bind uri
	if err := c.c.ShouldBindUri(obj); err != nil {
		return err
	}

	return nil
}
func (c *C) FormatRespOK(data interface{}) {
	c.c.JSON(http.StatusOK, generateResponse(data, http.StatusOK, "OK"))
}

func (c *C) FormatRespError(code int, message string) {
	c.c.JSON(code, generateResponse(nil, code, message))
}

func generateResponse(data interface{}, code int, message string) resp {
	return resp{
		Data: data,
		Status: stStatus{
			Code:    code,
			Message: message,
		},
	}
}

type resp struct {
	Data   any      `json:"data"`
	Status stStatus `json:"status"`
}

type stStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
