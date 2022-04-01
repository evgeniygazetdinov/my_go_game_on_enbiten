package main

import (
	"fmt"
	// "os"

	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	objs "github.com/SakoDroid/telego/objects"
	// "reflect"
)

const token string = "478565486:AAGyN1ey0LDTybgGlgdlZBeC3Bu73BMMxx0"

const relocate_group string = "-1001619080966"
const tractor_group string = "-1001212511273"

//The instance of the bot.
var bot *bt.Bot

func getUpdatesFromChannels(){
	fmt.Println("get updates")
}

func getPreferences(update *objs.Update) (int, int) {
	chatId := update.Message.Chat.Id
	messageId := update.Message.MessageId
	return chatId, messageId
}
//TODO find how getUpdates from channels

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
		// kip_up := <-*kiprChannel
		// fmt.Println(reflect.TypeOf(kip_up))
		// getUpdatesFromChannels()
		// if up.Message.Text == "rust" {
			_, err := bot.SendMessage(up.Message.Chat.Id, "f", "", up.Message.MessageId, false,false)
			
			if err != nil {
				fmt.Println(err)
			}

		}
	}
// }

	
