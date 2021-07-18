package commands

import (
	"botGo/checkers"
	"botGo/data"
	"os"
	"regexp"

	"github.com/bwmarrin/discordgo"
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
		if checkers.AdminCheck(s, m) {
			if userID != "" {
				if !checkers.CheckFile("data/" + checkers.CheckGuildId(s, m) + "/id/" + userID + ".json") {
					data.TxtTest(userID, s, m)
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
		if checkers.AdminCheck(s, m) {
			if userID != "" {
				if checkers.CheckFile("data/" + checkers.CheckGuildId(s, m) + "/id/" + userID + ".json") {
					os.Remove("data/" + checkers.CheckGuildId(s, m) + "/id/" + userID + ".json")
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
