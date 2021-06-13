package checkers

import (
	"github.com/bwmarrin/discordgo"
)

func AdminCheck(channelID string, userID string, s *discordgo.Session) bool {
	channelData, _ := s.Channel(channelID)
	guildID := channelData.GuildID
	guildMemberData, _ := s.GuildMember(guildID, userID)
	guildData, _ := s.Guild(guildID)

	var chAdmin bool
	chAdmin = false

	for _, z := range guildMemberData.Roles {

		var roleData *discordgo.Role

		for _, r := range guildData.Roles {
			if r.ID == z {
				roleData = r
			}
		}
		admin := roleData.Permissions & 0x00000008
		if admin == 8 {
			return true
		}
	}

	return chAdmin
}
