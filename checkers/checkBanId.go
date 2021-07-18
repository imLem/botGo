package checkers

import "github.com/bwmarrin/discordgo"

func CheckBanId(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	if CheckFile("data/" + CheckGuildId(s, m) + "/id/" + m.Author.ID + ".json") {
		return true
	} else {
		return false
	}
}
