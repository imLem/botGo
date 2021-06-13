package commands

import (
	"github.com/bwmarrin/discordgo"
	"regexp"
  "net/http"
	"botGo/checkers"
)

var validAva = regexp.MustCompile(`^[Ss][Pp][Oo]`)

func SpoilerHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}




	if checkers.CheckId(m.Author.ID) || validAva.MatchString(m.Content) {
		var urlFile string
		var fileName string
		var urlFile2 string
		for _, user2 := range m.Attachments {
			urlFile = user2.URL
			fileName = user2.Filename
		}
		for _, user3 := range m.Embeds {
			urlFile2 = user3.URL
		}
		if urlFile2 != "" {
			spoEmbed := "<@" + m.Author.ID + "> отправил медиа: " + "||" + urlFile2 + "||"
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			s.ChannelMessageSend(m.ChannelID, spoEmbed)
		}
		if m.Attachments != nil {
			//fmt.Println(urlFile)
			response, err := http.Get(urlFile) //use package "net/http"

			if err != nil {
				//fmt.Println(err)
				return
			}

			defer response.Body.Close()
			newName := "SPOILER_" + fileName
			menName := "<@" + m.Author.ID + "> отправил медиа:"
			//fmt.Println("Number of bytes copied to STDOUT:", n)
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			//s.ChannelFileSend(m.ChannelID, newName, response.Body)
			s.ChannelFileSendWithMessage(m.ChannelID, menName, newName, response.Body)
		}
	}
}
