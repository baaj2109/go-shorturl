package main

import "github.com/baaj2109/shorturl/router"

func main() {
	engine := router.InitRouter()
	panic(engine.Run())
}
