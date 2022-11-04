package usecase

import (
	"github.com/rs/zerolog/log"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type Request struct {
	Path string
}

type Response struct {
	Body   []byte
	Status int
}

func (svc *Service) ProcessRequest(req *Request) *Response {
	log.Debug().Str("path", req.Path).Msg("request received")
	return &Response{
		Body:   []byte("Hello there!"),
		Status: 200,
	}
}
