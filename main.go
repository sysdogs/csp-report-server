package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	log.Error().RawJSON("report", ctx.PostBody()).Send()
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Starting HTTP server at 8080")
	fasthttp.ListenAndServe(":8080", fastHTTPHandler)
}
