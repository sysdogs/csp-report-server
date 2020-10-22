package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.URI().Path()) == "/csp-report" {
		log.Error().RawJSON("message", ctx.PostBody()).RawJSON("url", ctx.URI().FullURI()).Send()
	} else {
		log.Info().RawJSON("url", ctx.URI().FullURI()).Send()
	}
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "timestamp"

	log.Info().Msg("Starting HTTP server at 8080")
	fasthttp.ListenAndServe(":8080", fastHTTPHandler)
}
