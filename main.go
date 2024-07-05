package main

import (
  "github.com/hsiang086/intel-fest/router"
)

func main() {
	router.Init()
	
	router.Routes()

	router.Server.ListenAndServe()
}
