package api

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/IvantimeKai/cron-job/global"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func SendEmailForVipUserEvery3Seconds(cr *cron.Cron) {
	fmt.Println(".....SendEmailForVipUserEvery3Seconds")

	_, err := cr.AddFunc("*/3 * * * * *", func() {
		log.Println("....run.....3 second")
	})

	if err != nil {
		global.GB_LOGGER.Error("cron not active", zap.Any("err", err))
	}
}

func GetUserInfoForApiGithubEvery5Seconds(cr *cron.Cron) {
	fmt.Println(".....GetUserInfoForApiGithubEvery5Seconds")

	_, err := cr.AddFunc("*/5 * * * * *", func() {
		rs, err := http.Get("https://api.github.com/users/IvantimeKai")

		if err != nil {
			global.GB_LOGGER.Error("user info not active", zap.Any("err", err))
			return
		}

		body, err := io.ReadAll(rs.Body)

		if err != nil {
			global.GB_LOGGER.Error("body not active", zap.Any("err", err))
			return
		}

		log.Println("github response", string(body))
	})

	if err != nil {
		global.GB_LOGGER.Error("cron not active", zap.Any("err", err))
	}
}
