package context

import (
	"encoding/json"
	"io"
	"net/http"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: context
 * @Version: ...
 * @Date: 2023-02-20 15:03:45
 */

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(obj interface{}) error {
	r := c.R
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, data interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)
	return err
}

func (c *Context) OkJson(data interface{}) error {
	return c.WriteJson(http.StatusOK, data)
}

func (c *Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}
func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		W: writer,
		R: request,
	}

}
