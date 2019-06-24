package main

import (
	"flag"
	"log"

	dg "github.com/bwmarrin/discordgo"
	"github.com/vharitonsky/iniflags"
)

var (
	discordBotToken = flag.String("secret", "", "Discord Bot Token")
)

func main() {
	iniflags.Parse()
	if discordBotToken == nil || *discordBotToken == "" {
		log.Fatalln("No token provided")
	}
	session, err := dg.New("Bot " + *discordBotToken)
	if err != nil {
		log.Fatalln("Failed to establish Discord session:" + err.Error())
	}

	handler := func(s *dg.Session, m *dg.MessageCreate) {

	}

	session.AddHandler(handler)
}
