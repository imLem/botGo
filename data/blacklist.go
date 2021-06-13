package data

import (
	"io/ioutil"
  "encoding/json"
)

type User struct {
	ID string
}

func TxtTest(id string) {
  data := User{
        ID: id,
        }
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("data/id/"+id+".json", file, 0644)
}
