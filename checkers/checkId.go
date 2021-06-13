package checkers


func CheckId(id string) bool {
	if CheckFile("data/id/" + id + ".json") {
		return true
	} else {
		return false
	}
}
