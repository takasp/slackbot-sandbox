package interactivemessages

import (
	"net/http"

	"github.com/nlopes/slack"
	"github.com/takasp/slackbot-sandbox/go-nlopes/common"
)

func ResponseMessage(w http.ResponseWriter, original slack.Message, value string) {
	original.Attachments[0].Actions = []slack.AttachmentAction{}
	original.Attachments[0].Fields = []slack.AttachmentField{
		{
			Value: value,
			Short: false,
		},
	}
	original.ResponseType = "in_channel"
	original.ReplaceOriginal = true

	common.SendMessage(w, &original.Msg)
}
