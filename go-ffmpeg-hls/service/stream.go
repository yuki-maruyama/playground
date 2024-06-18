package service

import (
	"context"

	"github.com/yuki-maruyama/playground/go-ffmpeg-hls/model"
)

type StreamService interface {
	ConvertStreamService(ctx context.Context, req model.StreamConvertRequest) (*model.StreamConvertResponse, error)
}

type streamService struct {
}

var _ StreamService = (*streamService)(nil)

func NewStreamService() *streamService {
	return &streamService{}
}

func (s *streamService) ConvertStreamService(ctx context.Context, req model.StreamConvertRequest) (*model.StreamConvertResponse, error) {
	return nil, nil
}
