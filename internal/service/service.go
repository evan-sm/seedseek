package service

import (
	"seedseek/internal/config"
	"seedseek/pkg/jackett"
	"seedseek/pkg/logger"
)

type Servicer interface {
}

type service struct {
	log     logger.Logger
	cfg     *config.Config
	jackett jackett.Jacketter
}

func New(log logger.Logger, cfg *config.Config, jackett jackett.Jacketter) Servicer {
	return &service{
		log:     log,
		cfg:     cfg,
		jackett: jackett,
	}
}
