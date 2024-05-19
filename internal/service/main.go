package service

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/Waramoto/hryvnia-svc/internal/config"
	"github.com/Waramoto/hryvnia-svc/internal/service/sender"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
}

func (s *service) run(cfg config.Config) error {
	s.log.Info("Service started")
	r := s.router(cfg)

	if err := s.copus.RegisterChi(r); err != nil {
		return fmt.Errorf("cop failed: %w", err)
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
	}
}

func Run(cfg config.Config) {
	ctx := context.Background()

	go sender.NewSender(cfg).Run(ctx)

	if err := newService(cfg).run(cfg); err != nil {
		panic(fmt.Errorf("failed to run service: %w", err))
	}
}
