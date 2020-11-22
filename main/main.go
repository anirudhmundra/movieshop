package main

import (
	"movieshop/router"
)

func main() {
	r := router.SetupRouter()
	r.Run()
}
