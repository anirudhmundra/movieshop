package main

import (
	"bookshop/router"
)

func main()  {
	r := router.SetupRouter()
	r.Run()
}
