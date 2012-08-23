package pushover

const message_max = 512
const url_max = 500
const url_title_max = 50

type Identity struct {
	token string
	user  string
}

type Message struct {
	text      string
	device    string
	title     string
	url       string
	url_title string
	priority  int
	timestamp int
}
