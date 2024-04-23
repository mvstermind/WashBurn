package handler

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func New(token string) {

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating session,", err)
		return
	}

	dg.AddHandler(MessageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("BOT IS WORKINNNNG!!!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "gen" {
		s.ChannelMessageSend(m.ChannelID, "WIELKA POLSKA GUROM")
	}
}
