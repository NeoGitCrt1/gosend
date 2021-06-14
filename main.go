package main

import (
	"flag"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
	"time"
	_ "time"

	gomail "gopkg.in/mail.v2"
)

// "cdtuhgrbalryenwl"
func main() {
	//from := flag.String("from", "test@d.a", "send from")
	to := flag.String("to", "test@d.c;test@d.d", "send to")
	num := flag.Int("num", 1, "how many to send")
	u := flag.String("u", "", "username: xxx@yahoo.com")
	p := flag.String("p", "", "generated app password")
	body := flag.String("body", "Hello geeks!!!", "send mail body")
	title := flag.String("title", "Hello geeks!!!", "send mail title")
	host := flag.String("hp", "smtp.mail.yahoo.com:587", "host+port")

	flag.Parse()
	toList := strings.Split(*to, ";")
	auth := smtp.PlainAuth("", *u, *p, strings.Split(*host, ":")[0])
	fmt.Println(auth)
	fmt.Println(toList)

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", *u)

	// Set E-Mail receivers
	m.SetHeader("To", *to)

	// Set E-Mail subject
	m.SetHeader("Subject", *title)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", *body)
	for i := 0; i < *num; i++ {

		// Settings for SMTP server
		intVar, _ := strconv.Atoi(strings.Split(*host, ":")[1])
		d := gomail.NewDialer(strings.Split(*host, ":")[0], intVar, *u, *p)

		// handling the errors
		if err := d.DialAndSend(m); err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("end")
}
