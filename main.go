package main

import "etri-sfpoc-cloud/router"

func main() {
	router.NewRouter().Run(":8080")
}
