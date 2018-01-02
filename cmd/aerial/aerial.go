package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/PonyvilleFM/aura/bot"
	"github.com/PonyvilleFM/aura/commands/source"
	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

type aerial struct {
	cs *bot.CommandSet
	s  *discordgo.Session
}

const (
	djonHelp  = ``
	djoffHelp = ``
	setupHelp = ``
)

func (a *aerial) Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := a.cs.Run(s, m.Message)
	if err != nil {
		log.Println(err)
	}
}

var (
	token             = os.Getenv("TOKEN")
	youtubeSpamRoomID = os.Getenv("DISCORD_YOUTUBESPAM_ROOMID")
	gClientID         = os.Getenv("GOOGLE_CLIENT_ID")
	gClientSecret     = os.Getenv("GOOGLE_CLIENT_SECRET")

	musicLinkRegex = regexp.MustCompile(`(.*)((http(s?):\/\/(www\.)?soundcloud.com\/.*)|(http(s?):\/\/(www\.)?youtube.com\/.*)|(http(s?):\/\/(www\.)?youtu.be\/.*))(.*)|(.*)http(s?):\/\/(www\.)?mixcloud.com\/.*`)
)

func main() {
	flag.Parse()
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	a := &aerial{
		cs: bot.NewCommandSet(),
		s:  dg,
	}

	a.cs.Prefix = ";"
	a.cs.AddCmd("np", "shows radio station statistics for Ponyville FM", bot.NoPermissions, stats)
	a.cs.AddCmd("stats", "shows radio station statistics for Ponyville FM", bot.NoPermissions, stats)
	a.cs.AddCmd("dj", "shows which DJ is up next on Ponyville FM", bot.NoPermissions, dj)
	a.cs.AddCmd("schedule", "shows the future radio schedule for Ponyville FM", bot.NoPermissions, schedule)
	a.cs.AddCmd("printerfact", "shows useful facts about printers", bot.NoPermissions, printerFact)
	a.cs.AddCmd("hipster", "hip me up fam", bot.NoPermissions, hipster)
	a.cs.AddCmd("source", "source code information", bot.NoPermissions, source.Source)
	a.cs.AddCmd("time", "shows the current bot time", bot.NoPermissions, curTime)
	a.cs.AddCmd("streams", "shows the different Ponyville FM stream links", bot.NoPermissions, streams)
	a.cs.AddCmd("derpi", "grabs a random **__safe__** image from Derpibooru with the given search results", bot.NoPermissions, derpi)

	dg.AddHandler(a.Handle)
	dg.AddHandler(pesterLink)
	dg.AddHandler(messageCreate)
	dg.AddHandler(imageMeEvent)

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("ready")

	<-make(chan struct{})
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Print message to stdout.
	fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)
}
