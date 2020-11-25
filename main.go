package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New("secret", "access")

	if err != nil {
		log.Fatalf("error linebot %s", err)
	}
	engine := gin.Default()
	engine.Any("/bot-hook", func(context *gin.Context) {
		events, err := bot.ParseRequest(context.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				context.Status(http.StatusBadRequest)
			} else {
				context.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
		}
		for _, event := range events {
			log.Printf("what is event %v %v\n", event, bot.GetProfile())
			if event.Type == linebot.EventTypeFollow {
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Hi nice to meet u")).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	})
	engine.Run(":4000")
}
