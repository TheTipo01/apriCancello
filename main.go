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
	// API key for the ESP8266
	key string
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
	key = cfg.apiKey

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

	// Start bot
	lit.Info("apriCancello is now running")
	b.Start()
}
