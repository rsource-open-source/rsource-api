package discord

// perms int: 35840
// TODO: Add rate limit

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rsource-open-source/rsource-api/utils"
)

func StartBot() {

	utils.LoadEnv()

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	dg.UserUpdateStatus("online")

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("bot online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Listen for api events and post a message to the discord channel
	// Then get that discord image link and return it

	// dg.ChannelMessageSend("944410841812459520", "hai")
	// dg.ChannelMessageSendComplex("944410841812459520", &discordgo.MessageSend{
	// 	Content: "hai",
	// 	Files: []*discordgo.File{

	// 	},
	// })
	dg.Close()
}
