package main

import (
	"apiservice/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
