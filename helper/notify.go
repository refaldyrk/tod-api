package helper

import (
	"context"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

func SendNotify(subj, text string) {
	telegramService, _ := telegram.New("5561567514:AAHtzrHCaVUTGVCmEiwFKUwO3pwqNa2dPpg")

	telegramService.AddReceivers(1993082483)
	notify.UseServices(telegramService)

	_ = notify.Send(
		context.Background(),
		subj,
		text,
	)
}
