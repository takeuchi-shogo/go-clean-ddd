package controllers

type Context interface {
	JSON(code int, i interface{}) error
	String(code int, s string) error
}
