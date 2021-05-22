package server

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/sky4access/extractor/internal/client"
)

type Service struct {
	logger *logrus.Entry
	client *client.Client
	ServiceOptions
}

type ServiceOptions struct {
	Debug bool
	Pretty bool
	Port string

}
func NewService(logger *logrus.Entry) *Service {
	p := new(Service)
	c, _ := client.NewClient("","","")
	p.client = c
	return p
}


func (s *Service) Run() error {
	return errors.New("not implemented")
}