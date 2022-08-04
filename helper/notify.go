package helper

import (
	"context"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

func SendNotify(subj, text string) {
	telegramService, _ := telegram.New("ID BOT TELEGRAM")

	telegramService.AddReceivers(123456)
	notify.UseServices(telegramService)

	_ = notify.Send(
		context.Background(),
		subj,
		text,
	)
}
