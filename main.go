package main

import (
	"fmt"
	// "os"

	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	"reflect"
)

const token string = "478565486:AAGyN1ey0LDTybgGlgdlZBeC3Bu73BMMxx0"

//The instance of the bot.
var bot *bt.Bot



func main() {
	up := cfg.DefaultUpdateConfigs()

	cf := cfg.BotConfigs{
		BotAPI: cfg.DefaultBotAPI,
		APIKey: token, UpdateConfigs: up,
		Webhook:        false,
		LogFileAddress: cfg.DefaultLogFile,
	}

	var err error

	//Creating the bot using the created configs
	bot, err = bt.NewBot(&cf)

	if err == nil {

		err = bot.Run()

		if err == nil {
			start()
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}
func start() {
	//Register the channel
	messageChannel, _ := bot.AdvancedMode().RegisterChannel("", "message")

	for {
		//Wait for updates
		up := <-*messageChannel
		fmt.Println(reflect.TypeOf(up))

		if up.Message.Text == "rust" {
			_, err := bot.SendMessage(up.Message.Chat.Id, "hi to you too, send me a location", "", up.Message.MessageId, false,false)
			if err != nil {
				fmt.Println(err)
			}

		}
	}
}

	
// func getPreferences(update *objects.Update) (string, string) {
// 	chatId := update.Message.Chat.Id
// 	messageId := update.Message.MessageId
// 	return chatId, messageId
// }