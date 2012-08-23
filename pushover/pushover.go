package pushover

import (
	"log"
	"net/http"
	"net/url"
	//"time"
)

func loginfo(message string) {
	if Verbose {
		message = "[+] pushover " + message
		log.Println(message)
	}
}

func logerr(message string) {
	if Verbose {
		message = "[!] pushover " + message
		log.Println(message)
	}
}

func member_too_long(mtype string, mlength int, maxlen int) {
	if Verbose {
		log.Printf("[!] pushover warning: %s length of %d chars exceeds "+
			"maximum allowed %s length\n    of %d chars.\n",
			mtype, mlength, mtype, maxlen)
	}
}

// returns a boolean indicating whether the message was valid. if the
// message was invalid, the offending struct member(s) was/were
// truncated.
func validate_message(message Message) (Message, bool) {
	valid := true

	if len(message.token) == 0 {
		logerr("missing authentication token.")
		valid = false
	}

	if len(message.user) == 0 {
		logerr("missing user key.")
		valid = false
	}

	if len(message.text) == 0 {
		logerr("missing message.")
		valid = false
	}

	message_len := len(message.text) + len(message.title)
	if message_len > message_max {
		member_too_long("message", message_len, message_max)
		message.text = message.text[:message_max-len(message.title)]
		valid = false
	}

	if len(message.url) > url_max {
		member_too_long("URL", len(message.url), url_max)
		valid = false
	}

	if len(message.url_title) > url_title_max {
		member_too_long("URL title", len(message.url_title),
			url_title_max)
		valid = false
	}

	return message, valid
}

func get_body(message Message) (url.Values, bool) {
	body := url.Values{}
	valid := true

	body.Add("token", message.token)
	body.Add("user", message.user)
	body.Add("message", message.text)

	if len(message.device) > 0 {
		body.Add("device", message.device)
	}

	if len(message.title) > 0 {
		body.Add("title", message.title)
	}

	if len(message.url) > 0 {
		body.Add("url", message.url)
	}

	if len(message.url_title) > 0 {
		body.Add("url_title", message.url_title)
	}

	if len(message.priority) > 0 {
		body.Add("priority", message.priority)
	}

	if len(message.timestamp) > 0 {
		body.Add("timestamp", message.timestamp)
	}

	return body, valid
}

func notify(message Message) bool {
	_, valid := validate_message(message)
	if !valid {
		logerr("invalid message")
	}

	body, valid := get_body(message)
	if !valid {
		logerr("invalid message.")
		return valid
	}

	loginfo("sending message...")
	resp, err := http.PostForm(api_url, body)
	if err != nil {
		logerr("POST request failed.")
		return false
	} else {
		defer resp.Body.Close()
	}

	loginfo("POST request success.")
	if resp.StatusCode != 200 {
		logerr("server returned " + resp.Status + ".")
		return false
	}
	return true
}

func Authenticate(token string, user string) Identity {
	return Identity{token, user}
}

func Notify(identity Identity, message string) bool {
	msg := Message{identity.Token, identity.User, message, "", "", "", "",
		"0", ""}
	if !err {
		log.Println("[!] error creating message.")
		return false
	}

	return notify(msg)
}

func Notify_titled(identity Identity, message string, title string) bool {
	msg := Message{identity.Token, identity.User, message, "", title, "", "",
		"0", ""}
	return notify(msg)
}

func Notify_device(identity Identity, message string, device string) bool {
	msg := Message{identity.Token, identity.User, message, device, "", "", "",
		"0", ""}
	return notify(msg)
}
