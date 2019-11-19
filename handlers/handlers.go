package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/martip07/minecraftbox-api/procs"
	"github.com/martip07/minecraftbox-api/structs"

	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	respHome := structs.GeneralMessage{
		Message: "API minecraftbox v1",
	}
	jsonBody, err := json.Marshal(respHome)

	if err != nil {
		ctx.Error(" json marshal fail", 500)
		return

	}
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetStatusCode(200)
	ctx.Response.SetBody(jsonBody)
	return
}

func NotFoundHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(405)
	return
}

func StreamingHandler(ctx *fasthttp.RequestCtx) {
	if ctx.QueryArgs().Has("game") && ctx.QueryArgs().Has("lang") {
		gameARG := string(ctx.QueryArgs().Peek("game"))
		langARG := string(ctx.QueryArgs().Peek("lang"))
		respMessage := procs.TwitchProc(gameARG, langARG)

		jsonBody, err := json.Marshal(respMessage)

		if err != nil {
			fmt.Println(err)
			ctx.Error(" json marshal fail", 500)
			return

		}
		ctx.SetContentType("application/json; charset=utf-8")
		ctx.Response.Header.Set("Cache-Control", "public, max-age=600")
		ctx.Response.Header.Set("Server", "go-minecraftbox")
		ctx.SetStatusCode(200)
		ctx.Response.SetBody(jsonBody)
		return
	}
	respMessage := structs.GeneralMessage{
		Message: "Nothing around here",
	}
	jsonBody, err := json.Marshal(respMessage)

	if err != nil {
		fmt.Println(err)
		ctx.Error(" json marshal fail", 500)
		return

	}
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.Response.Header.Set("Cache-Control", "public, max-age=600")
	ctx.Response.Header.Set("Server", "go-minecraftbox")
	ctx.SetStatusCode(200)
	ctx.Response.SetBody(jsonBody)
	return

}
