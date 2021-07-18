package data

import (
	"botGo/checkers"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
)

type User struct {
	ID string
}

func TxtTest(id string, s *discordgo.Session, m *discordgo.MessageCreate) {
	data := User{
		ID: id,
	}
	file, _ := json.MarshalIndent(data, "", " ")
	if !checkers.CheckFile("data/" + checkers.CheckGuildId(s, m) + "/id") {
		err := os.MkdirAll("data/"+checkers.CheckGuildId(s, m)+"/id", 0777)
		if err != nil {
			panic(err)
		}
	}
	_ = ioutil.WriteFile("data/"+checkers.CheckGuildId(s, m)+"/id/"+id+".json", file, 0644)
}
