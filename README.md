## gopush
### a Go client library for Pushover notifications

This is my first real Go program, so *caveat emptor*.

`gopush` provides an interface to [Superblock's](https://superblock.net)
[Pushover](https://pushover.net).

To use it, use the import `"bitbucket.org/kisom/gopush/pushover"`.

Example usage:

    package main
    
    import (
    	"fmt"
    	"bitbucket.org/kisom/gopush/pushover"
    	"os"
    	"path/filepath"
    )
    
    func main() {
    	if len(os.Args) < 3 {
    		fmt.Printf("usage: %s api_key user_key\n", filepath.Base(os.Args[0]))
    		os.Exit(1)
    	}
    
    	identity := pushover.Authenticate(
    		os.Args[1],
    		os.Args[2],
    	)
    
    	sent := pushover.Notify(identity,"testing gopush")
    	if !sent {
    		fmt.Println("[!] notification failed.")
    		os.Exit(1)
    	}
    }
    
## Notification functions:

### Authenticate
`Authenticate` returns an Identity struct. To create one, call the Authentication
function with the API token and user key:

    identity := pushover.Authenticate(token, user_key)

### Notify
`Notify` is the most basic notification function. Its signature is:
    
    Notify(Identity, string) bool

`Notify` returns `true` if the message was sent successfully, and `false`
otherwise.

### Notify\_titled
`Notify_titled` is used to send a notification with a custom title. Its 
signature is:

    Notify_titled(Identity, string, string) bool

The first string is the message, the second string is the title. It returns the
same as `Notify`.

### Notify\_device
`Notify_device` is used to send a notification to a specific device. Its 
signature is:

    Notify_titled(Identity, string, string) bool

The first string is the message, the second string is the device name. It returns
the same as `Notify`.
