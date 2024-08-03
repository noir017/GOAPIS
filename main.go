package main

import (
	"github.com/noir017/goapis/app"
	"github.com/noir017/goapis/routers"
)

func main() {
	app.Init()
	routers.StartService()
}
