// Handles the imap interactions

package imap

import (
	"log"
	"os"
	//	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	//    "strconv"
)

type Account struct {
	name      string
	email     string
	mailboxes []Mailbox
}
type Mailbox struct {
	// config    config // store config here
	name     string
	children []Mailbox
	mail     []Mail
}

type Mail struct {
	fromName  string
	fromEmail string
	subject   string
	body      string
}

// Just a proof of concept
func Test() {

	var user, pass, host string

	user = os.Getenv("imap_user")
	pass = os.Getenv("imap_pass")
	host = os.Getenv("imap_host")

	log.Println("Connecting to " + host + " with " + user + ":::" + pass + "...")

	account := Account{name: host, email: user}

	log.Println("Created account object for " + account.name)
	// Connect to server
	c, err := client.DialTLS(host+":993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(user, pass); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")

	// TODO : Add caching to this, currently loading on every execution from the server
	// ideally, this entire datastructure can be stored locally in json, and then
	// refreshed in the background once it's displayed
	//
	// also perhaps add a cron CLI flag you can add to crontab to precache new mail offline in the cache
	//

	// TODO: add proper support for nested mailboxes, currently they just use dot-notation to seperate levels
	for m := range mailboxes {
		var mailbox Mailbox = Mailbox{ name:m.Name }
		account.mailboxes = append(account.mailboxes, mailbox)
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", mbox.Flags)

	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only substract if the result is > 0
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []string{imap.EnvelopeMsgAttr}, messages)
	}()

	log.Println("Last 4 messages:")
	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")

}
