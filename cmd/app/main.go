package main

import (
	"github.com/daniwira/multifinance/router"
)

func main() {
	r, err := router.SetupRouter()
	if err != nil {
		panic(err)
	}

	r.Run()
}
