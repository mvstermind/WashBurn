package handler

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mvstermind/halset/generator"
)

//var token string = "your token"

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
	if m.Content == "generate minor" || m.Content == "gen min" || m.Content == ":3" {
		gned := generator.Gen{}
		bpmstr := strconv.Itoa(gned.GetBpm()) // had to do this cuz bot accepts only string
		wholeMessage := fmt.Sprintf("BPM: %v", bpmstr)
		s.ChannelMessageSend(m.ChannelID, wholeMessage)
		k, v := Midi(1)
		s.ChannelFileSend(m.ChannelID, k, v)

	}
}
