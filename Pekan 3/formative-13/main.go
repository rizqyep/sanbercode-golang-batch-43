package main

import "formative-13/routers"

func main() {
	PORT := ":8080"
	routers.StartServer().Run(PORT)
}
