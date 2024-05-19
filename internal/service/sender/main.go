package sender

import (
	"context"
	"net/smtp"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"

	"github.com/Waramoto/hryvnia-svc/internal/config"
	"github.com/Waramoto/hryvnia-svc/internal/data"
	pg "github.com/Waramoto/hryvnia-svc/internal/data/postgres"
)

const senderRunnerName = "sender"

type Sender struct {
	db     data.SubscribersQ
	logger *logan.Entry
	config config.Config
	auth   smtp.Auth
}

func NewSender(cfg config.Config) *Sender {
	emailConfig := cfg.SenderConfig().Email

	return &Sender{
		db:     pg.NewDB(cfg.DB()).Subscribers(),
		logger: cfg.Log(),
		config: cfg,
		auth: smtp.PlainAuth(
			emailConfig.Identity,
			emailConfig.From,
			emailConfig.Password,
			emailConfig.Host,
		),
	}
}

func (s *Sender) Run(ctx context.Context) {
	s.logger.Info("Starting sender...")
	runnerConfig := s.config.SenderConfig().Runner
	running.WithBackOff(ctx,
		s.logger,
		senderRunnerName,
		s.Send,
		runnerConfig.NormalPeriod,
		runnerConfig.MinAbnormalPeriod,
		runnerConfig.MaxAbnormalPeriod,
	)
}
