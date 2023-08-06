package bot

import (
	"context"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var btn = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonURL("Поиск", "https://google.com"),
	),
)

func (b *Bot) CmdFormat(ctx context.Context, u tg.Update) (err error) {
	b.bot.Send(tg.NewChatAction(u.Message.Chat.ID, tg.ChatTyping))

	var text strings.Builder

	m := tg.NewMessage(u.Message.From.ID, text.String())
	// m.ParseMode = tg.ModeMarkdown
	m.ParseMode = tg.ModeHTML
	b.bot.Send(m)

	return nil
}
