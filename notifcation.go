package main

import (
	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	objs "github.com/SakoDroid/telego/objects"
	"fmt"
	"strconv"
 )
 

const token string = "478565486:AAGyN1ey0LDTybgGlgdlZBeC3Bu73BMMxx0"

 func main(){
    up := cfg.DefaultUpdateConfigs()
    
    cf := cfg.BotConfigs{BotAPI: cfg.DefaultBotAPI, APIKey: token, UpdateConfigs: up, Webhook: false, LogFileAddress: cfg.DefaultLogFile}

    bot, err := bt.NewBot(&cf)

    if err == nil{

        t := err == bot.Run()
		fmt.Println(t)
        if err == nil{
            go start(bot)
        }
    }
 }

 func start(bot *bt.Bot){

     //The general update channel.
     updateChannel := bot.GetUpdateChannel()

    //Adding a handler. Everytime the bot receives message "hi" in a private chat, it will respond "hi to you too".

    bot.AddHandler("hi",func(u *objs.Update) {

        //Register channel for receiving messages from this chat.
		cc, _ := bot.AdvancedMode().RegisterChannel(strconv.Itoa(u.Message.Chat.Id), "message")

        //Sends back a message
		_, err := bot.SendMessage(u.Message.Chat.Id, "hi to you too, send me a location", "", u.Message.MessageId, false,false)
		if err != nil {
			fmt.Println(err)
		}

        //Waits for an update from this chat
		up := <-*cc

        //Sends back the received location
		_, err = bot.SendLocation(up.Message.Chat.Id, false,false, up.Message.Location.Latitude, up.Message.Location.Longitude, up.Message.Location.HorizontalAccuracy, up.Message.MessageId)


		if err != nil {
			fmt.Println(err)
		}
	},"private")

    //Monitores any other update. (Updates that don't contain text message "hi" in a private chat)
     for {
         update := <- *updateChannel
		fmt.Println(update.Message.Text)
        //Some processing on the update
     }
 }