package main

import (
	"github.com/hsiang086/intel-fest/database"
	"github.com/hsiang086/intel-fest/router"
	"github.com/hsiang086/intel-fest/router/api"
)

func main() {
	database.Init()

	api.Init()

	router.Init()

	router.Routes()

	router.Server.ListenAndServe()
}
