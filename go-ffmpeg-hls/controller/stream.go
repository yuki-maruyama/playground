package controller

import (
	"net/http"

	"github.com/yuki-maruyama/playground/go-ffmpeg-hls/service"
)

type streamController struct {
	service service.StreamService
}

func NewStreamController(s service.StreamService) *streamController {
	return &streamController{service: s}
}

func (c *streamController) ConvertVideoHandler(w http.ResponseWriter, r *http.Request) {

}
