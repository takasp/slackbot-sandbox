package bento

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/takasp/slackbot-sandbox/go-nlopes/interactivemessages"

	"github.com/nlopes/slack"
	"github.com/takasp/slackbot-sandbox/go-nlopes/common"
)

func ActionBento(w http.ResponseWriter, message slack.AttachmentActionCallback, action slack.AttachmentAction) {
	switch action.Name {
	case "bentoActionSelect":
		value := action.SelectedOptions[0].Value

		// Overwrite original drop down message.
		originalMessage := message.OriginalMessage
		originalMessage.Attachments[0].Text = fmt.Sprintf("%s を注文してよろしいですか?", strings.Title(value))
		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{
			{
				Name:  "bentoActionStart",
				Text:  "Yes",
				Type:  "button",
				Value: "start",
				Style: "primary",
			},
			{
				Name:  "bentoActionCancel",
				Text:  "No",
				Type:  "button",
				Style: "danger",
			},
		}
		originalMessage.ResponseType = "in_channel"
		originalMessage.ReplaceOriginal = true
		common.SendMessage(w, &originalMessage.Msg)
		return
	case "bentoActionStart":
		title := ":ok: <@" + message.User.ID + "> の注文を受け付けました。"
		interactivemessages.ResponseMessage(w, message.OriginalMessage, title)
		return
	case "bentoActionCancel":
		title := ":x: <@" + message.User.ID + "> の注文をキャンセルしました。"
		interactivemessages.ResponseMessage(w, message.OriginalMessage, title)
		return
	default:
		log.Printf("[ERROR] ]Invalid action was submitted: %s", action.Name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
