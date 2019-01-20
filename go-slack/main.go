package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nlopes/slack"

	"github.com/takasp/slackbot-sandbox/go-nlopes/handler"

	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	Host              string `envconfig:"HOST" default:"localhost"`
	Port              string `envconfig:"PORT" default:"3000"`
	VerificationToken string `envconfig:"VERIFICATION_TOKEN" required:"true"`
	OAuthAccessToken  string `envconfig:"OAUTH_ACCESS_TOKEN" required:"true"`
	ChannelID         string `envconfig:"CHANNEL_ID" required:"true"`
}

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

func _main() error {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		return fmt.Errorf("[ERROR] Failed to process env var: %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "healthy")
	})

	slackClient := slack.New(env.OAuthAccessToken)
	slackClient.SetDebug(true)

	http.Handle("/interaction", handler.InteractionHandler{
		SlackClient:       slackClient,
		VerificationToken: env.VerificationToken,
		ChannelID:         env.ChannelID,
	})

	http.Handle("/slash", handler.SlashHandler{
		SlackClient:       slackClient,
		VerificationToken: env.VerificationToken,
		ChannelID:         env.ChannelID,
	})

	fmt.Printf("[INFO] Server listening on %s port %s", env.Host, env.Port)
	if err := http.ListenAndServe(env.Host+":"+env.Port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	return nil
}
