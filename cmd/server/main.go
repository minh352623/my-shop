package main

import "ecom/internal/routers"

func main() {
	r := routers.InitRouter()

	r.Run() // listen and serve on 0.0.0.0:8080
}
