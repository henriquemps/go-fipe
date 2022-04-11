package main

import "github.com/joho/godotenv"

func init() {
	godotenv.Load()

	registerRoutes()
}

func main() {
	serverUp()
}
