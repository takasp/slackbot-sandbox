package bento

import (
	"net/http"

	"github.com/nlopes/slack"
	"github.com/takasp/slackbot-sandbox/go-nlopes/slashcommands"
)

func OrderBento(w http.ResponseWriter, slashCommand slack.SlashCommand, channelID string) {

	slashcommands.ChannelPermission(w, slashCommand, channelID)

	attachment := slack.Attachment{
		Fallback:   "Select bento",
		Text:       "注文する弁当を選択してください。",
		CallbackID: "bento_selection",
		Actions: []slack.AttachmentAction{
			{
				Name: "bentoActionSelect",
				Type: "select",
				Options: []slack.AttachmentActionOption{
					{Text: "のり弁当", Value: "のり弁当"},
					{Text: "唐揚げ弁当", Value: "唐揚げ弁当"},
					{Text: "幕の内弁当", Value: "幕の内弁当"},
					{Text: "銀しゃけ弁当", Value: "銀しゃけ弁当"},
					{Text: "しょうが焼き弁当", Value: "しょうが焼き弁当"},
				},
			},
			{
				Name:  "bentoActionCancel",
				Text:  "Cancel",
				Type:  "button",
				Style: "danger",
			},
		},
	}

	var attachments []slack.Attachment

	msg := &slack.Msg{ResponseType: "in_channel", Attachments: append(attachments, attachment)}
	slashcommands.SendMessage(w, msg)
	return
}
