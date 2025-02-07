package config

import (
	"strconv"
)

const APP_VERSION = "1.0.4"

var (
	BOT_TOKEN             = Env("BOT_TOKEN")
	ADMIN_ID, _           = strconv.ParseInt(Env("ADMIN_ID"), 10, 64)
	CHANNEL_ID            = Env("CHANNEL_ID")
	CHAPPIE_SERVER_URL    = Env("CHAPPIE_SERVER_URL")
	CHAPPIE_SERVER_BEARER = Env("CHAPPIE_SERVER_BEARER")
	WAPP_TOKEN            = Env("WAPP_TOKEN")
	WAPP_JID              = Env("WAPP_JID")
	WAPP_SERVER_URL       = Env("WAPP_SERVER_URL")
	WAPP_ENABLE           = func() bool {
		if val := Env("WAPP_ENABLE"); val != "" {
			parsed, err := strconv.ParseBool(val)
			if err == nil {
				return parsed
			}
		}
		return false
	}()
	ENABLE_CRON = func() bool {
		if val := Env("ENABLE_CRON"); val != "" {
			parsed, err := strconv.ParseBool(val)
			if err == nil {
				return parsed
			}
		}
		return false
	}()
)
