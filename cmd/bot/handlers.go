package main

import (
	"context"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/dialog"
)

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
    //warnHandler(ctx,b,update)
    //dialogHandler(ctx,b,update)
    switch {
        case update.Message != nil  :
            warnHandler(ctx,b,update)
        case update.CallbackQuery != nil  :
	        dialogHandler(ctx,b,update)
    }
}

func helloHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Hello, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "*",
		ParseMode: models.ParseModeMarkdown,
	})
}

func warnHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Este bot no es interactive,solo acepta comands",
	})
}

func replyHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}

func dialogHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	p := dialog.New(dialogNodes)
	p.Show(ctx, b, strconv.Itoa(update.Message.Chat.ID), "start")
}

func dialogInlineHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	p := dialog.New(dialogNodes, dialog.Inline())
	p.Show(ctx, b, strconv.Itoa(update.Message.Chat.ID), "start")
}

func adoptHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "adoptar, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "*",
		ParseMode: models.ParseModeMarkdown,
	})
}

func giveInAdoptionHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "dar en adopcion, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "*",
		ParseMode: models.ParseModeMarkdown,
	})
}
