package bot

import (
	"fmt"

	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/moabukar/discord-bot-go/config"
)

var (
	BotID string
	goBot *discordgo.Session
)

func Start() {

	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// if m.Author.ID == BotID {
	// 	return
	// }

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pongg!")
	}

	if m.Content == "salaam" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "walaukum salaam!")
	}

	if m.Content == "assalamu alaykum" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "walaukum salaam wrwb!")
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "hello"
	if strings.ToLower(m.Content) == "hello" {
		// Send a message with nohello.com
		s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	}

	if strings.ToLower(m.Content) == "hi" {
		// Send a message with nohello.com
		s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	}

	if strings.Contains(m.Content, s.State.User.Mention()) {
		s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	}

	if strings.ToLower(m.Content) == "nh" {
		// Send a message with nohello.com
		s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	}
}
