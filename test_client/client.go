package main

import (
	"fmt"
	"gopush/pushover"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s api_key user_key\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	fmt.Println("[+] testing pushover library.")
	fmt.Println("[+] identifying...")
	identity := pushover.Authenticate(
		os.Args[1],
		os.Args[2],
	)
	fmt.Println("[+] identification complete.")
	fmt.Printf("[+] user key: %s\n", identity.User)

	sent := pushover.Notify(identity, "testing gopush")
	if !sent {
		fmt.Println("[!] notification failed.")
		os.Exit(1)
	}

	sent = pushover.Notify_titled(identity, "testing gopush", "gopush test")
	if !sent {
		fmt.Println("[!] notification failed.")
		os.Exit(1)
	}

	sent = pushover.Notify_device(identity, "testing gopush", "galaxynexus")
	if !sent {
		fmt.Println("[!] notification failed.")
		os.Exit(1)
	}
}
