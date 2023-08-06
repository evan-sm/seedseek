package bot

import (
	"context"
	"fmt"
	"seedseek/internal/indexer"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) CmdTest(ctx context.Context, u tg.Update) (err error) {
	b.bot.Send(tg.NewChatAction(u.Message.Chat.ID, tg.ChatTyping))

	items, err := b.idx.Fetch(ctx, "", indexer.SortingSeed)
	if err != nil {
		return err
	}

	if len(items) >= 5 {
		items = items[:5]
	}

	var text strings.Builder

	for i, item := range items {
		text.WriteString(fmt.Sprint(i+1, ". "))
		text.WriteString(fmt.Sprintf("<a href=\"%s\">%s</a>\n", item.Comments, item.Title[:41]))

		text.WriteString(fmt.Sprintf("📁%.2d GB 🔼 %d 🔽 %d ⏬️ %d ", bytesToGB(item.Size), item.Seeders, item.Peers, item.Grabs))
		text.WriteString(fmt.Sprintf("<a href=\"%s\">💾 DL</a>\n\n", item.Link))
	}

	m := tg.NewMessage(u.Message.From.ID, text.String())
	m.ParseMode = tg.ModeHTML
	b.bot.Send(m)

	return nil
}
