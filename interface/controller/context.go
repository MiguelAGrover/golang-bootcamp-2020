package controller

// Context : interface serves a carrier for incoming data in Bind and Param, and Json for the response
type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(i string) string
}
