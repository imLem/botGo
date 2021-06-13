package commands

import (
	"botGo/data"
	"github.com/bwmarrin/discordgo"
	"os"
	"regexp"
  "botGo/checkers"
)

var blockCall = regexp.MustCompile(`^[Ff][Ii][Rr][Ee]`)
var unblockCall = regexp.MustCompile(`^[Ff][Rr][Ee][Ee]`)

func AddBlockHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	var userID string

	if blockCall.MatchString(m.Content) {
		for _, user := range m.Mentions {
			userID = user.ID
		}
		if checkers.AdminCheck(m.ChannelID, m.Author.ID, s) {
			if userID != "" {
				if !checkers.CheckFile("data/id/" + userID + ".json") {
					data.TxtTest(userID)
					s.ChannelMessageSend(m.ChannelID, "<@"+userID+"> занесен в список")
				} else {
					s.ChannelMessageSend(m.ChannelID, "Пользователь уже в списке")
				}
			} else {
				s.ChannelMessageSend(m.ChannelID, "Никто не указан")
			}
		} else {
      s.ChannelMessageSend(m.ChannelID, "Нет прав")
    }
	}

	if unblockCall.MatchString(m.Content) {
		for _, user := range m.Mentions {
			userID = user.ID
		}
		if checkers.AdminCheck(m.ChannelID, m.Author.ID, s) {
			if userID != "" {
				if checkers.CheckFile("data/id/" + userID + ".json") {
					os.Remove("data/id/" + userID + ".json")
					s.ChannelMessageSend(m.ChannelID, "<@"+userID+"> убран из списка")
				} else {
					s.ChannelMessageSend(m.ChannelID, "Пользователя нет в списке")
				}

			} else {
				s.ChannelMessageSend(m.ChannelID, "Никто не указан")
			}

		} else {
      s.ChannelMessageSend(m.ChannelID, "Нет прав")
    }
	}

}
