package checkers

import (
	"github.com/bwmarrin/discordgo"
)

func AdminCheck(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	channelData, _ := s.Channel(m.ChannelID)
	guildID := channelData.GuildID
	guildMemberData, _ := s.GuildMember(guildID, m.Author.ID)
	guildData, _ := s.Guild(guildID)

	for _, z := range guildMemberData.Roles {
		for _, r := range guildData.Roles {
			if r.ID == z {
				if (r.Permissions & 0x00000008) != 0 {
					return true
				}
			}
		}
	}
	return false
}
