package bot

import (
	"context"
	"log"
	"net/http"

	"seedseek/internal/config"
	"seedseek/internal/indexer"
	"seedseek/pkg/logger"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Boter interface {
	Run(ctx context.Context) error
	Close() error
}

type Bot struct {
	log logger.Logger
	cfg *config.Config
	bot *tg.BotAPI
	idx indexer.Indexer
}

func New(ctx context.Context, cfg *config.Config, log logger.Logger, idx indexer.Indexer) (Boter, error) {
	bot, err := tg.NewBotAPI(cfg.BotToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.InfoContext(ctx, "Authorized on account %s", bot.Self.UserName)

	return &Bot{
		log: log,
		cfg: cfg,
		bot: bot,
		idx: idx,
	}, nil
}

func (b *Bot) Run(ctx context.Context) error {
	wh, err := tg.NewWebhook(b.cfg.WebHookURL + b.bot.Token)
	if err != nil {
		return err
	}

	_, err = b.bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	updates := b.bot.ListenForWebhook("/" + b.bot.Token)
	go http.ListenAndServe("0.0.0.0:8443", nil)

	for u := range updates {
		if u.Message != nil {
			b.log.InfoContext(ctx, "[%s:%s] %s", u.Message.Chat.Title, u.Message.From.UserName, u.Message.Text)
		}

		if u.Message == nil {
			continue
		}

		if !u.Message.IsCommand() {
			continue
		}

		switch u.Message.Command() {
		case CmdStart:
			err = b.CmdStart(ctx, u)
			if err != nil {
				b.log.ErrorContext(ctx, err.Error())
				continue
			}
		case CmdHelp:
			err = b.CmdHelp(ctx, u)
			if err != nil {
				b.log.ErrorContext(ctx, err.Error())
				continue
			}
		case CmdPing:
			err = b.CmdPing(ctx, u)
			if err != nil {
				b.log.ErrorContext(ctx, err.Error())
				continue
			}
		// case CmdTest:
		// 	err = b.CmdTest(ctx, u)
		// 	if err != nil {
		// 		b.log.ErrorContext(ctx, err.Error())
		// 		continue
		// 	}
		case CmdGet:
			err = b.CmdGet(ctx, u)
			if err != nil {
				b.log.ErrorContext(ctx, err.Error())
				continue
			}
			// case CmdFormat:
			// 	err = b.CmdFormat(ctx, u)
			// 	if err != nil {
			// 		b.log.ErrorContext(ctx, err.Error())
			// 		continue
			// 	}

		}

	}

	return nil
}

func (b *Bot) Close() error {
	return nil
}

func bytesToGB(bytes uint) uint {
	gigabytes := bytes / (1024 * 1024 * 1024)
	return gigabytes
}
