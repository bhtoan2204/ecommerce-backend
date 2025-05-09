package handler

import "gateway/package/wrapper"

type RequestHandler interface {
	Handle(ctx *wrapper.Context) (interface{}, error)
}
