package botwaygo

import (
	"bytes"
	"errors"

	"github.com/abdfnx/botway/constants"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func GetBotInfo(value string) string {
	viper.SetConfigType("yaml")

	viper.ReadConfig(bytes.NewBuffer(constants.BotConfig))

	return viper.GetString(value)
}

func GetToken() string {
	if GetBotInfo("bot.lang") != "go" {
		panic(errors.New("ERROR: Your Bot language is not Go"))
	} else {
		data := gjson.Get(string(constants.BotwayConfig), "botway.bots."+GetBotInfo("bot.name")+".bot_token").String()

		return data
	}
}

func GetAppId() string {
	if GetBotInfo("bot.lang") != "go" {
		panic(errors.New("ERROR: Your Bot language is not Go"))
	} else {
		id := "bot_app_id"

		if GetBotInfo("bot.type") == "slack" {
			id = "bot_app_token"
		}

		data := gjson.Get(string(constants.BotwayConfig), "botway.bots."+GetBotInfo("bot.name")+"."+id).String()

		return data
	}
}

func GetSigningSecret() string {
	if GetBotInfo("bot.lang") != "go" {
		panic(errors.New("ERROR: Your Bot language is not Go"))
	} else {
		data := gjson.Get(string(constants.BotwayConfig), "botway.bots."+GetBotInfo("bot.name")+".signing_secret").String()

		return data
	}
}

func GetGuildId(serverName string) string {
	if GetBotInfo("bot.lang") != "go" {
		panic(errors.New("ERROR: Your Bot language is not Go"))
	} else if GetBotInfo("bot.type") != "discord" {
		panic(errors.New("ERROR: This function/feature is only working with discord bots"))
	} else {
		data := gjson.Get(string(constants.BotwayConfig), "botway.bots."+GetBotInfo("bot.name")+".guilds."+serverName+".server_id").String()

		return data
	}
}
