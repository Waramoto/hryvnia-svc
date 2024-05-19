package config

import (
	"fmt"
	"time"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type SenderConfig struct {
	Period time.Duration `fig:"period,required"`
	Email  EmailConfig   `fig:"email,required"`
	Runner RunnerConfig  `fig:"runner,required"`
}

type SenderConfiger interface {
	SenderConfig() *SenderConfig
}

func NewSenderConfiger(getter kv.Getter) SenderConfiger {
	return &senderConfig{
		getter: getter,
	}
}

type senderConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func (s *senderConfig) SenderConfig() *SenderConfig {
	return s.once.Do(func() any {
		configMap := kv.MustGetStringMap(s.getter, "sender")
		var cfg SenderConfig
		err := figure.Out(&cfg).From(configMap).Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out sender config: %w", err))
		}

		return &cfg
	}).(*SenderConfig)
}
