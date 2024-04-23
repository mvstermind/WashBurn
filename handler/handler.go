package handler

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/mvstermind/halset/generator"
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
	if m.Content == "generate" || m.Content == "gen" || m.Content == ":3" {
		gned := generator.Gen{}
		bpmstr := strconv.Itoa(gned.GetBpm()) // had to do this cuz bot accepts only string
		chordsStr := strings.Join(gned.GetChords(), " ")
		if m.Content == ":3" {

			wholeMessage := fmt.Sprintf("How lovely :3\nBPM: %v\nKey: %v\nChords: %v\n", bpmstr, gned.GetKey(), chordsStr)
			s.ChannelMessageSend(m.ChannelID, wholeMessage)
		} else {
			wholeMessage := fmt.Sprintf("BPM: %v\nKey: %v\nChords: %v\n", bpmstr, gned.GetKey(), chordsStr)
			s.ChannelMessageSend(m.ChannelID, wholeMessage)
		}
	}

}
