package main

import (
	"fmt"
	"log"

	"github.com/martip07/minecraftboxapi/handlers"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Println("Howdy from minecraftbox API v1")
	r := router.New()
	r.GET("/api/", handlers.Index)
	r.GET("/api/profiler/twitch", handlers.AuthHandler)
	r.GET("/api/streaming/twitch", handlers.StreamingHandler)
	log.Fatal(fasthttp.ListenAndServe(":8081", r.Handler))
}
