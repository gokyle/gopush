package pushover

import (
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
func validate_message(message Message) bool {
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

	return valid
}

func basic_message(message string) Message {
	if length(message) > message_max {
	}
	return Message{message, "", "", "", "", 0,
		time.Now().UTC().Unix(),
	}
}

func send_basic(message Message, identity Identity) int {

}
