package main

import (
	"dependencies/router"
)

func main() {
	// Generate all routes
	r := router.Generate()

	r.Run()
}
