package commands

import (
	"C"
	"botGo/checkers"
	"net/http"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var validAva = regexp.MustCompile(`^[Ss][Pp][Oo]`)

func SpoilerHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if checkers.CheckBanId(s, m) || validAva.MatchString(m.Content) {
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
			response, err := http.Get(urlFile)
			if err != nil {
				return
			}
			defer response.Body.Close()
			newName := "SPOILER_" + fileName
			menName := "<@" + m.Author.ID + "> отправил медиа:"
			s.ChannelMessageDelete(m.ChannelID, m.ID)
			s.ChannelFileSendWithMessage(m.ChannelID, menName, newName, response.Body)
		}
	}
}
