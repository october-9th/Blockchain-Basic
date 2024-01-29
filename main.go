package main

import "blockchain.com/m/api"

func main() {
	router := api.CreateRouter()

	router.Run()
}
