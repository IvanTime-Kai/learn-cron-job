package global

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var (
	GB_CRON   *cron.Cron
	GB_LOGGER *zap.Logger
)
