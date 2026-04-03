package registry

import (
	"github.com/IvantimeKai/cron-job/api"
	"github.com/IvantimeKai/cron-job/global"
)

func RegistryApiRunCronjob() {
	api.SendEmailForVipUserEvery3Seconds(global.GB_CRON)
	api.GetUserInfoForApiGithubEvery5Seconds(global.GB_CRON)

	global.GB_CRON.Start()
}
