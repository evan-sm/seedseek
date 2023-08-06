package bot

import (
	"context"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const ()

func (b *Bot) CmdStart(ctx context.Context, u tg.Update) error {
	b.bot.Send(tg.NewChatAction(u.Message.Chat.ID, tg.ChatTyping))

	b.bot.Send(tg.NewMessage(u.Message.From.ID, "🤖 Hi! I'm SeedSeek! Try /get command."))

	return nil
}
