package controllers

type Context interface {
	FormValue(name string) string
	JSON(code int, i interface{}) error
	QueryParam(name string) string
	Param(name string) string
	String(code int, s string) error
}
