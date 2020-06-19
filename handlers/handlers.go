package handlers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/martip07/minecraftboxapi/procs"
	"github.com/martip07/minecraftboxapi/structs"

	"gopkg.in/h2non/gentleman.v2"

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

func AuthHandler(ctx *fasthttp.RequestCtx) {
	type tokenData struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `token_type:"id"`
	}

	cli := gentleman.New()
	cli.Method("POST")
	clientId := os.Getenv("CLIENTID")
	clientSecret := os.Getenv("CLIENTSECRET")
	uriBase := "https://id.twitch.tv/oauth2/token?client_id=" + clientId + "&client_secret=" + clientSecret + "&grant_type=client_credentials"
	res, err := cli.Request().URL(uriBase).Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
	}
	json := &tokenData{}
	res.JSON(json)
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
