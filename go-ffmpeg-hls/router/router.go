package router

import (
	"net/http"

	"github.com/yuki-maruyama/playground/go-ffmpeg-hls/api/middleware"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	middlewares := []func(http.Handler) http.Handler{
		middleware.LoggingMiddleware,
	}

	streamDirFunc := http.FileServer(http.Dir("./stream"))
	mux.Handle("/stream/", streamDirFunc)
	mux.HandleFunc("GET /")

	return applyMiddleware(mux, middlewares...)
}

func applyMiddleware(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
