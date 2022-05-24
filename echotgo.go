package main

import (
	"github.com/wasmcloud/actor-tinygo"
	httpserver "github.com/wasmcloud/interfaces/httpserver/tinygo"
)

func main() {
	me := Echotgo{}
	actor.RegisterHandlers(httpserver.HttpServerHandler(&me), actor.ActorHandler(&me))
}

type Echotgo struct{}

func (e *Echotgo) HandleRequest(ctx *actor.Context, req httpserver.HttpRequest) (*httpserver.HttpResponse, error) {
	r := httpserver.HttpResponse{
		StatusCode: 200,
		Header:     make(httpserver.HeaderMap, 0),
		Body:       []byte("hello world"),
	}
	return &r, nil
}

func (e *Echotgo) HealthRequest(ctx *actor.Context, arg actor.HealthCheckRequest) (*actor.HealthCheckResponse, error) {
	var r actor.HealthCheckResponse
	r.Healthy = true
	return &r, nil
}
