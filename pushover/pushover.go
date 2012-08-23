package pushover

import (
	"encodings/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func authenticate(token string, user string) Identity {
	return Identity{token, user}
}

func member_too_long(mtype string, mlength int, maxlen int) {
	log.Printf("[!] warning: %s length of %d chars exceeds maximum allowed "+
		"%s length\n    of %d chars.\n", mtype, length, mtype, maxlen)
}

// returns a boolean indicating whether the message was valid. if the
// message was invalid, the offending struct member(s) was/were
// truncated.
func validate_message(message Message) (Message, bool) {
	valid := true
	message_len := length(message.text) + length(message.title)
	if message_len > message_max {
		member_too_long("message", message_len, message_max)
		message = message[:message_max-length(message.title)]
		valid = false
	}

	if length(message.url) > url_max {
		member_too_long("URL", length(message.url), url_max)
		valid = false
	}

	if length(message.url_title) > url_max {
		member_too_long("URL title", length(message.url_title),
			max_url_title)
		valid = false
	}

	return message, valid
}

func basic_message(message string, identity Identity) (Message, bool) {
	msg := Message{identity.token, identity.user, message, "", "", "", "",
		0, time.Now().UTC().Unix()}
	var valid bool

	msg, valid = validate_message(msg)
	return msg, valid
}

func notify(message Message, identity Identity) bool {
	log.Println("[+] encoding message to JSON")
	json_message, json_err := json.Marshall(message)
	if json_err {
		log.Println("[!] error encoding to JSON.")
		return json_err
	}

	log.Printf("[-] body: '%s'\n", json_message)

	log.Println("[+] sending message...")
	resp, err := http.Post(api_url, "application/json", json_message)
	if err != nil {
		log.Printf("[!] POST request failed with error %s.\n", err)
		return err
	} else {
		defer resp.Body.Close()
	}

	log.Println("[+] POST request success.")
	if resp.StatusCode != 200 {
		log.Printf("[!] server returned %s.\n", resp.Status)
		return false
	}
}
