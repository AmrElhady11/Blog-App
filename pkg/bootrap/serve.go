package bootrap

import (
	"blog/internal/moduls"
	"blog/pkg/config"
	"blog/pkg/routing"
)

func Serve() {
	config.Set()
	routing.Init()
	moduls.Routes(routing.GetRouter())
	routing.Serve()
}
