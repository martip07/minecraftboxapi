package main

import (
	"fmt"
	"log"

	"github.com/martip07/minecraftbox-api/handlers"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Println("Howdy from demo main")
	// TwitchProcEs()
	r := router.New()
	r.GET("/", handlers.Index)
	r.GET("/streaming/twitch", handlers.StreamingHandler)
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}