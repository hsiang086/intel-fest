package main

import (
	_ "github.com/hsiang086/intel-fest/database"
	"github.com/hsiang086/intel-fest/router"
)

func main() {
	router.Routes()

	router.Server.ListenAndServe()
}
