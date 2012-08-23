package pushover

import (
	"encoding/json"
	"log"
	"net/http"
        "strings"
	"time"
)

func authenticate(token string, user string) Identity {
	return Identity{token, user}
}

func member_too_long(mtype string, mlength int, maxlen int) {
	log.Printf("[!] warning: %s length of %d chars exceeds maximum allowed "+
		"%s length\n    of %d chars.\n", mtype, mlength, mtype, maxlen)
}

// returns a boolean indicating whether the message was valid. if the
// message was invalid, the offending struct member(s) was/were
// truncated.
func validate_message(message Message) (Message, bool) {
	valid := true
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

func basic_message(message string, identity Identity) (Message, bool) {
	msg := Message{identity.token, identity.user, message, "", "", "", "",
		0, int(time.Now().UTC().Unix())}
	var valid bool

	msg, valid = validate_message(msg)
	return msg, valid
}

func notify(message Message, identity Identity) bool {
	log.Println("[+] encoding message to JSON")
	json_message, json_err := json.Marshal(message)
	if json_err != nil {
		log.Println("[!] error encoding to JSON.")
		return false
	}

        message_body := strings.NewReader(string(json_message))
	log.Printf("[-] body: '%s'\n", message_body)

	log.Println("[+] sending message...")
	resp, err := http.Post(api_url, "application/json", message_body)
	if err != nil {
		log.Printf("[!] POST request failed with error %s.\n", err)
		return false
	} else {
		defer resp.Body.Close()
	}

	log.Println("[+] POST request success.")
	if resp.StatusCode != 200 {
		log.Printf("[!] server returned %s.\n", resp.Status)
		return false
	}
        return true
}
