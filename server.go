package main

import "about-me/handler"

func main() {
	handler.SetUpServer().Run("localhost:9000")
}
