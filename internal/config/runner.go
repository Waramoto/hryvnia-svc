package config

import "time"

type RunnerConfig struct {
	NormalPeriod      time.Duration `fig:"normal_period,required"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period,required"`
	MaxAbnormalPeriod time.Duration `fig:"max_abnormal_period,required"`
}
