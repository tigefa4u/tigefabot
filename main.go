package main

import (
	"github.com/go-chat-bot/bot/irc"
	_ "github.com/go-chat-bot/plugins/godoc"
	_ "github.com/go-chat-bot/plugins/catfacts"
	_ "github.com/go-chat-bot/plugins/catgif"
	_ "github.com/go-chat-bot/plugins/gif"
	_ "github.com/go-chat-bot/plugins/chucknorris"
	_ "github.com/go-chat-bot/plugins/uptime"
	_ "github.com/go-chat-bot/plugins/url"
	_ "github.com/go-chat-bot/plugins/web"
	_ "github.com/go-chat-bot/plugins/encoding"
	_ "github.com/go-chat-bot/plugins/puppet"
	_ "github.com/go-chat-bot/plugins/guid"
	_ "github.com/go-chat-bot/plugins/treta"
	_ "github.com/go-chat-bot/plugins/crypto"

	// Import all the commands you wish to use
	"os"
	"strings"
)

func main() {
	irc.Run(&irc.Config{
		Server:   os.Getenv("IRC_SERVER"),
		Channels: strings.Split(os.Getenv("IRC_CHANNELS"), ","),
		User:     os.Getenv("IRC_USER"),
		Nick:     os.Getenv("IRC_NICK"),
		Password: os.Getenv("IRC_PASSWORD"),
		UseTLS:   true,
		Debug:    os.Getenv("DEBUG") != ""})
}
