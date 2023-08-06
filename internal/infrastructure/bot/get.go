package bot

import (
	"context"
	"fmt"
	"seedseek/internal/indexer"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	lineLimit     = 44
	lineLimitMono = 34
)

func (b *Bot) CmdGet(ctx context.Context, u tg.Update) (err error) {
	b.bot.Send(tg.NewChatAction(u.Message.Chat.ID, tg.ChatTyping))

	var sort string

	switch u.Message.CommandArguments() {
	default:
		sort = indexer.SortingSeed
	case "gain":
		sort = indexer.SortingGain
	}

	items, err := b.idx.Fetch(ctx, "", sort, 5)
	if err != nil {
		return err
	}

	var text strings.Builder
	var title string

	for i, item := range items {
		title = item.Title
		if len(title) > lineLimit-3 {
			title = title[:lineLimit-3]
		}

		text.WriteString(fmt.Sprint(i+1, ". "))
		text.WriteString(fmt.Sprintf("<a href=\"%s\">%s</a>\n", item.Details, title))

		text.WriteString(fmt.Sprintf("📁%.2d GB 🔼 %d 🔽 %d ⏬️ %d ", bytesToGB(item.Size), item.Seeders, item.Peers, item.Grabs))
		text.WriteString(fmt.Sprintf("📈 %.2f ", item.Gain))
		text.WriteString(fmt.Sprintf("<a href=\"%s\">💾 DL</a>\n\n", item.Link))
	}

	m := tg.NewMessage(u.Message.From.ID, text.String())
	m.ParseMode = tg.ModeHTML
	m.DisableWebPagePreview = true

	b.bot.Send(m)

	return nil
}
