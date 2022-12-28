package bot

import (
	"fmt"

	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/moabukar/no-hello-bot/config"
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

	if m.Content == "salaam!" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "walaykum salaam!")
	}

	if m.Content == "assalamu alaykum" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "walaykum salaam wrwb!")
	}

	// If the message is "hello"
	// if strings.ToLower(m.Content) == "hello" {
	// 	// Send a message with nohello.com
	// 	s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	// }

	switch strings.ToLower(m.Content) {
	// If the message is "hello"
	case "hi", "hello", "wag1", "ello", "lo", "h3llo", "hiya", "hi there", "yo", "hey", "hell0", "h3y", "h3ll0":
		// Send a message with nohello.com
		_, _ = s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	default:
		// empty
	}

	switch strings.ToLower(m.Content) {
	// If the message is "hello"
	case "salaam", "salaams":
		// Send a message with nohello.com
		_, _ = s.ChannelMessageSend(m.ChannelID, "Walaykum salaam!")
	default:
		// empty
	}

	switch strings.ToLower(m.Content) {
	// If the message is "hello"
	case "inshaallah", "insha'allah", "in sha allah", "inshallah", "ان شاء الله", "ia":
		// Send a message with nohello.com
		_, _ = s.ChannelMessageSend(m.ChannelID, "inshaAllah")
	default:
		// empty
	}

	switch strings.ToLower(m.Content) {
	case "asc":
		_, _ = s.ChannelMessageSend(m.ChannelID, "wcs wrwb!")
	default:
		// empty
	}

	switch strings.ToLower(m.Content) {
	case "haha":
		_, _ = s.ChannelMessageSend(m.ChannelID, "😂")
	default:
		// empty
	}

	switch strings.ToLower(m.Content) {
	case "jzk khair", "jazakallahu khair", "جزاك الله خيرا", "jazakallahu khairan", "jazaakallah Khair", "jazakallahu khayr", "jazaakallahu khayr", "jazakallahu khayran", "jazakAllahu khairaa", "jazakallah":
		_, _ = s.ChannelMessageSend(m.ChannelID, "wa iyyaaka")
	default:
		// empty
	}

	switch strings.ToLower(m.Content) {
	case "jazakumallahu khayran", "جزاكم الله خيرا", "جزاكم الله خير", "jazakumallahu khair", "jazakumallahu khayr", "jzk":
		_, _ = s.ChannelMessageSend(m.ChannelID, "wa iyyaakum")
	default:
		// empty
	}
	switch strings.ToLower(m.Content) {
	case "السلام عليكم":
		_, _ = s.ChannelMessageSend(m.ChannelID, "و عليكم السلام - Please ask your question too!")
	default:
		// empty
	}

	switch strings.ToLower(m.Content) {
	case "wow", "wow!", "woww", "wowww":
		_, _ = s.ChannelMessageSend(m.ChannelID, "WOW!!!")
	default:
		// empty
	}

	switch strings.ToLower(m.Content) {
	case "alh":
		_, _ = s.ChannelMessageSend(m.ChannelID, "Alhamdulillah")
	default:
		// empty
	}

	// if strings.ToLower(m.Content) == "hi" {
	// 	// Send a message with nohello.com
	// 	s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	// }

	if strings.Contains(m.Content, s.State.User.Mention()) {
		// Send a message with nohello.com when tagged.
		s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	}

	if strings.ToLower(m.Content) == "nh" {
		// Send a message with nohello.com when "nh" is used
		s.ChannelMessageSend(m.ChannelID, "https://nohello.com")
	}

	if strings.ToLower(m.Content) == "salaam" {
		// Send a prompt when a salaam is written
		s.ChannelMessageSend(m.ChannelID, "Walaykum salaam! How can we help 🙂 If you've got a query, make sure to put in as much detail as possible!")
	}

	if strings.ToLower(m.Content) == "salaams" {
		// Send a prompt when a salaam is written
		s.ChannelMessageSend(m.ChannelID, "Walaykum salaam! How can we help 🙂 If you've got a query, make sure to put in as much detail as possible!")
	}

	if strings.ToLower(m.Content) == "salam" {
		// Salaam variation 2.0
		s.ChannelMessageSend(m.ChannelID, "Walaykum salaam! How can we help 🙂 If you've got a query, make sure to put in as much detail as possible!")
	}

	if strings.ToLower(m.Content) == "salams" {
		// Salaam variation 2.0
		s.ChannelMessageSend(m.ChannelID, "Walaykum salaam! How can we help 🙂 If you've got a query, make sure to put in as much detail as possible!")
	}

	// if strings.ToLower(m.Content) == "assalamu alaykum" {
	// 	// Salaam variation 3.0
	// 	s.ChannelMessageSend(m.ChannelID, "walaukum salaam wrwb!")
	// }

	if strings.ToLower(m.Content) == "nc" {
		// No Context variation
		s.ChannelMessageSend(m.ChannelID, "Please give more context to your question!! Outline the issue at hand; giving enough detail and background for someone to help")
	}

	// if strings.ToLower(m.Content) == "jdi" {
	// 	// No Context variation
	// 	s.ChannelMessageSend(m.ChannelID, "just do it!!!")
	// }

	if strings.ToLower(m.Content) == "should i apply?" {
		// No Context variation
		s.ChannelMessageSend(m.ChannelID, "just apply!!!")
	}

	if strings.ToLower(m.Content) == "sia" {
		// No Context variation
		s.ChannelMessageSend(m.ChannelID, "just apply!!!")
	}

	if strings.ToLower(m.Content) == "ty" {
		// No Context variation
		s.ChannelMessageSend(m.ChannelID, "No problem!")
	}

	if strings.ToLower(m.Content) == "badq" {
		// No Context variation
		s.ChannelMessageSend(m.ChannelID, "https://discord.com/channels/844170468398202890/1050209853529915412")
	}

	if strings.ToLower(m.Content) == "badq1" {
		// No Context variation
		s.ChannelMessageSend(m.ChannelID, ".")
	}

	// str := "should i apply?"
	// matched, err := regexp.MatchString(`.*should i apply.*`, str)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if matched && str == "should i apply?" {
	// 	s.ChannelMessageSend(m.ChannelID, "just apply!!!")
	// }

	// str := "should i apply?"
	// matched, err := regexp.MatchString(`.*should i apply.*`, str)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// switch {
	// case matched && str == "should i apply?":
	// 	s.ChannelMessageSend(m.ChannelID, "just apply!!!")
	// default:

	switch strings.ToLower(m.Content) {
	case "jdi":
		_, _ = s.ChannelMessageSend(m.ChannelID, "just do it")
	default:
		// empty
	}

}
