package main
import (

	"fmt"
	// "os"
    
	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
//	objs "github.com/SakoDroid/telego/objects"
	"reflect"
)

// const token string = "478565486:AAGyN1ey0LDTybgGlgdlZBeC3Bu73BMMxx0"
//The instance of the bot.
var bot *bt.Bot

func getPreferences(update *objects.Update){
	chatId := update.Message.Chat.Id
	messageId := update.Message.MessageId
	return chatId, messageId
}

func main() {
	up := cfg.DefaultUpdateConfigs()
    
	cf := cfg.BotConfigs{
      BotAPI: cfg.DefaultBotAPI, 
      APIKey: token, UpdateConfigs: up, 
      Webhook: false, 
      LogFileAddress: cfg.DefaultLogFile,
    }
    
    var err error
	
    //Creating the bot using the created configs
    bot, err = bt.NewBot(&cf)
	
    if err == nil {
    
    	err = bot.Run()
        
        if err == nil{
			start()
        }else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
    
}
func start(){
	//Register the channel
	messageChannel, _ := bot.AdvancedMode().RegisterChannel("", "message")
    
    for {
    
    	//Wait for updates
		up := <- *messageChannel
        
		if(up.Message.Text	== "rust"){
			
		}
    }
}