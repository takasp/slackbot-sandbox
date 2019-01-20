package slashcommands

import (
	"encoding/json"
	"net/http"

	"github.com/nlopes/slack"
)

func SendMessage(w http.ResponseWriter, msg *slack.Msg) {
	b, err := json.Marshal(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func ChannelPermission(w http.ResponseWriter, slashCommand slack.SlashCommand, channelID string) {
	if slashCommand.ChannelID != channelID {
		params := &slack.Msg{Text: "このチャンネルでは使用できません。"}
		b, err := json.Marshal(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}
