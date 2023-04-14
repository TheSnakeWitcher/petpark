package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

type Bot struct { 
    *bot.Bot
}

func main() {
    fmt.Println("execution start")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	
	opts := []bot.Option{
		bot.WithDebug(),
		bot.WithDefaultHandler(dialogHandler),
		//bot.WithCallbackQueryDataHandler("/adopt",bot.MatchTypeExact,adoptHandler),
		//bot.WithCallbackQueryDataHandler("/giveInAdoption",bot.MatchTypeExact, giveInAdoptionHandler),
	}
	b, err := bot.New(os.Getenv("BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}


	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData,"/adopt", bot.MatchTypeExact, adoptHandler)
	//b.RegisterHandler(bot.HandlerTypeCallbackQueryData,"/giveInAdoption", bot.MatchTypeExact, giveInAdoptionHandler)

	//b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, helloHandler)
	//b.RegisterHandler(bot.HandlerTypeMessageText,"/adoptar",bot.MatchTypeExact,adoptHandler)
	//b.RegisterHandler(bot.HandlerTypeMessageText,"/reportar",bot.MatchTypeExact,reportHandler)


	//go b.StartWebhook(ctx)
	//http.ListenAndServe(":2000", b.WebhookHandler())
	b.Start(ctx)
    fmt.Println("execution end")
}
