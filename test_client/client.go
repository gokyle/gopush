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

        msg, valid := pushover.Basic_message("testing gopush", identity)
        if !valid {
            fmt.Println("[!] invalid message.")
            os.Exit(1)
        } 

        sent := pushover.Notify(msg)
        if !sent {
            fmt.Println("[!] notification failed.")
            os.Exit(1)
        }
}
