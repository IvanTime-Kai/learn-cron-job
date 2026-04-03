package main

import (
	"github.com/IvantimeKai/cron-job/global"
	"github.com/IvantimeKai/cron-job/initialize"
	"github.com/IvantimeKai/cron-job/registry"
)

func main() {
	global.GB_CRON = initialize.InitCron()
	global.GB_LOGGER = initialize.InitLogger()

	registry.RegistryApiRunCronjob()

	select {
		
	}
}
