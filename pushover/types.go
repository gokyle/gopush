package pushover

const message_max = 512
const url_max = 500
const url_title_max = 50
const api_url = "https://api.pushover.net/1/messages.json"

var Verbose = true

type Identity struct {
	Token string
	User  string
}

type Message struct {
	token     string
	user      string
	text      string
	device    string
	title     string
	url       string
	url_title string
	priority  string
	timestamp string
}
