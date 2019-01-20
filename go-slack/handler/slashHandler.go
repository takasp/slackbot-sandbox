package handler

import (
	"log"
	"net/http"

	"github.com/takasp/slackbot-sandbox/go-nlopes/slashcommands/bento"

	"github.com/nlopes/slack"
)

type SlashHandler struct {
	SlackClient       *slack.Client
	VerificationToken string
	ChannelID         string
}

func (h SlashHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("[ERROR] Invalid method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	slashCommand, err := slack.SlashCommandParse(r)
	if err != nil {
		log.Printf("[ERROR] Failed to parse slash command from slack: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !slashCommand.ValidateToken(h.VerificationToken) {
		log.Printf("[ERROR] Invalid token: %s", slashCommand.Token)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch slashCommand.Command {
	case "/bento":
		bento.OrderBento(w, slashCommand, h.ChannelID)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
