package main

import (
	"github.com/abdulkarim1422/BloodsApp/initializers"
	"github.com/abdulkarim1422/BloodsApp/lib"
	"github.com/abdulkarim1422/BloodsApp/routers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()

	lib.SetupLogOutput()
}

func main() {
	routers.AllRouters()
}
