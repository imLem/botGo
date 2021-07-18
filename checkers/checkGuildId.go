package checkers

import (
	"github.com/bwmarrin/discordgo"
)

func CheckGuildId(s *discordgo.Session, m *discordgo.MessageCreate) string {
	channelData, _ := s.Channel(m.ChannelID)
	return channelData.GuildID
}
