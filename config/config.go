package config

import (
	"strconv"
)

const APP_VERSION = "1.1.0"

var (
	BOT_TOKEN             = Env("BOT_TOKEN")
	ADMIN_ID, _           = strconv.ParseInt(Env("ADMIN_ID"), 10, 64)
	CHANNEL_ID            = Env("CHANNEL_ID")
	CHAPPIE_SERVER_URL    = Env("CHAPPIE_SERVER_URL")
	CHAPPIE_SERVER_BEARER = Env("CHAPPIE_SERVER_BEARER")
	WAPP_TOKEN            = Env("WAPP_TOKEN")
	WAPP_JID              = Env("WAPP_JID")
	X_API_KEY             = Env("X_API_KEY")
	X_URL                 = Env("X_URL")
	WAPP_SERVER_URL       = Env("WAPP_SERVER_URL")
	WAPP_ENABLE           = parseBoolEnv("WAPP_ENABLE")
	ENABLE_CRON           = parseBoolEnv("ENABLE_CRON")
)
