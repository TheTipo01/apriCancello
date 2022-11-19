package main

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func start(c tele.Context) error {
	msg := fmt.Sprintf("Il tuo ID è: %d", c.Sender().ID)

	if _, ok := whitelist[c.Sender().ID]; ok {
		msg += "\nSei nella whitelist."
	} else {
		msg += "\nNon sei nella whitelist."
	}

	return c.Send(msg)
}

func apri(c tele.Context) error {
	if _, ok := whitelist[c.Sender().ID]; ok {
		err := apertura()
		if err != nil {
			return c.Send("Errore nell'apertura: " + err.Error())
		} else {
			return c.Send("Apertura effettuata.")
		}
	} else {
		return c.Send("Non sei nella whitelist.")
	}
}
