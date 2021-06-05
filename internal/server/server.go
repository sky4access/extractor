package server

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/sky4access/extractor/internal/esvclient"
	"time"
)

type Service struct {
	logger *logrus.Entry
	client *esvclient.ESVClient
	ServiceOptions
}

type ServiceOptions struct {
	Debug  bool
	Pretty bool
	Port   string
}

func NewService(logger *logrus.Entry) *Service {
	p := new(Service)
	c, _ := esvclient.NewClient(2 * time.Second)
	p.client = c
	return p
}

func (s *Service) Run() error {
	return errors.New("not implemented")
}
