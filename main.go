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
	// Subscribed users
	whitelist map[int64]bool
	// Endpoint of the ESP8266
	endpoint string
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
	endpoint = cfg.Endpoint

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
	// Create bot
	b, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		lit.Error(err.Error())
		return
	}

	b.Handle("/start", start)
	b.Handle("/apri", apri)

	b.Handle(tele.OnText, apri)
	b.Handle(tele.OnAudio, apri)
	b.Handle(tele.OnDocument, apri)
	b.Handle(tele.OnPhoto, apri)
	b.Handle(tele.OnSticker, apri)
	b.Handle(tele.OnVideo, apri)
	b.Handle(tele.OnVoice, apri)
	b.Handle(tele.OnVideoNote, apri)

	// Keyboard
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnApri := menu.Text("Apri")

	b.Handle(&btnApri, apri)

	// Start bot
	lit.Info("apriCancello is now running")
	b.Start()
}
