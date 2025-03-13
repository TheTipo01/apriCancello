package main

import (
	"github.com/bwmarrin/lit"
	"github.com/kkyr/fig"
	tele "gopkg.in/telebot.v3"
	"strings"
	"time"
)

var (
	// Telegram token
	token string
	// Home Assistant token
	HAToken string
	// Subscribed users
	whitelist map[int64]bool
	// Endpoints of things to open
	endpoints map[string]endpoint
	// Menu
	menu = &tele.ReplyMarkup{ResizeKeyboard: true}
)

func init() {
	lit.LogLevel = lit.LogError

	var cfg config
	err := fig.Load(&cfg, fig.File("config.yml"))
	if err != nil {
		lit.Error(err.Error())
		return
	}

	// Config file found
	token = cfg.Token
	HAToken = cfg.HAToken

	// Set endpoints
	endpoints = make(map[string]endpoint, len(cfg.Endpoints))
	for _, e := range cfg.Endpoints {
		endpoints[e.Name] = e
	}

	whitelist = make(map[int64]bool, len(cfg.IDs))
	for _, id := range cfg.IDs {
		whitelist[id] = true
	}

	// Set lit.LogLevel to the given value
	switch strings.ToLower(cfg.LogLevel) {
	case "logwarning", "warning":
		lit.LogLevel = lit.LogWarning

	case "loginformational", "informational":
		lit.LogLevel = lit.LogInformational

	case "logdebug", "debug":
		lit.LogLevel = lit.LogDebug
	}
}

func main() {
	var err error
	var b *tele.Bot

	for {
		// Create bot
		b, err = tele.NewBot(tele.Settings{
			Token:  token,
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		})
		if err == nil {
			break
		} else {
			lit.Info("Can't connect to Telegram, retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		}
	}

	b.Handle("/start", start)
	b.Handle("/apri", apri)

	rows := make([]tele.Row, 0, len(endpoints))

	for k := range endpoints {
		btn := menu.Text(k)
		rows = append(rows, menu.Row(btn))
		b.Handle(&btn, apri)
	}

	// Keyboard
	menu.Reply(
		rows...,
	)

	// Start bot
	lit.Info("apriCancello is now running")
	b.Start()
}
