package main

import (
	"go-find-me/pkg/routes"
	"log"
	"os"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var (
	output = log.New(os.Stdout, "", 0)
)

func Logger(req fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		begin := time.Now()
		req(ctx)
		end := time.Now()
		output.Printf("[%v] %v | %s %s - %v - %v | %s",
			end.Format("2006/01/02 - 15:04:05"),
			ctx.RemoteAddr(),
			ctx.Method(),
			ctx.RequestURI(),
			ctx.Response.Header.StatusCode(),
			end.Sub(begin),
			ctx.UserAgent(),
		)
	})
}

func main() {
	r := router.New()

	routes.DeviceRoutes(r)
	routes.CheckerRoutes(r)

	log.Fatal(fasthttp.ListenAndServe(":8080", Logger(r.Handler)))
}
