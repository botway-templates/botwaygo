package botwaygo

import (
	"bytes"
	"errors"

	"github.com/abdfnx/botway/constants"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func GetBotInfo(value string) string {
	botConfig := viper.New()

	botConfig.SetConfigType("yaml")

	botConfig.ReadConfig(bytes.NewBuffer(constants.BotConfig))

	return botConfig.GetString(value)
}

func GetToken() string {
	data := gjson.Get(string(constants.BotwayConfig), "botway.bots."+GetBotInfo("bot.name")+".bot_token").String()

	return data
}

func GetAppId() string {
	id := "bot_app_id"

	if GetBotInfo("bot.type") == "slack" {
		id = "bot_app_token"
	}

	data := gjson.Get(string(constants.BotwayConfig), "botway.bots."+GetBotInfo("bot.name")+"."+id).String()

	return data
}

func GetSecret() string {
	value := ""

	if GetBotInfo("bot.type") == "slack" {
		value = "signing_secret"
	} else if GetBotInfo("bot.type") == "twitch" {
		value = "bot_client_secret"
	}

	data := gjson.Get(string(constants.BotwayConfig), "botway.bots."+GetBotInfo("bot.name")+"."+value).String()

	return data
}

func GetGuildId(serverName string) string {
	if GetBotInfo("bot.type") != "discord" {
		panic(errors.New("ERROR: This function/feature is only working with discord bots"))
	} else {
		data := gjson.Get(string(constants.BotwayConfig), "botway.bots."+GetBotInfo("bot.name")+".guilds."+serverName+".server_id").String()

		return data
	}
}
