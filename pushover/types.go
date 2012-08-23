package pushover

const message_max = 512
const url_max = 500
const url_title_max = 50
const api_url = "https://api.pushover.net/1/messages.json"

type Identity struct {
	Token string
	User  string
}

type Message struct {
	Token     string
	User      string
	Text      string
	Device    string
	Title     string
	Url       string
	Url_title string
	Priority  int
	Timestamp int
}
