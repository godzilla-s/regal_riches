package service

import "github.com/godzilla-s/regal-riches/pkg/model"

type Config struct {
	DBConfig *model.Config
	Url      string
}
