package main

import (
	"github.com/hsiang086/intel-fest/database"
	"github.com/hsiang086/intel-fest/router"
)

func main() {
	database.Init()

	router.Init()

	router.Routes()

	router.Server.ListenAndServe()
}
