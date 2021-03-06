package commands

import (
	"botGo/checkers"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var avatarCall = regexp.MustCompile(`^[Aa][Vv][Aa]`)

const sizeAvatar = "2048"

func AvatarHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	var userID string
	if avatarCall.MatchString(m.Content) {
		for _, user := range m.Mentions {
			userID = user.ID
		}
		if checkers.CheckBanId(s, m) {
			s.ChannelMessageSend(m.ChannelID, "Доступ ограничен")
		} else {
			userObj, _ := s.User(userID)
			if userObj != nil {
				s.ChannelMessageSend(m.ChannelID, userObj.AvatarURL(sizeAvatar))
			} else {
				s.ChannelMessageSend(m.ChannelID, m.Author.AvatarURL(sizeAvatar))
			}
		}
	}
}
