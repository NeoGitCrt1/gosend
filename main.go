package main

import (
	"flag"
	"fmt"
	"net/smtp"
	"strings"
	"time"
	_ "time"
)

// "cdtuhgrbalryenwl"
func main() {
	//from := flag.String("from", "test@d.a", "send from")
	toList := strings.Split(*(flag.String("to", "test@d.c;test@d.d", "send to")), ";")
	num := flag.Int("num", 1, "how many to send")
	u := flag.String("u", "", "username: xxx@yahoo.com")
	p := flag.String("p", "", "generated app password")
	body := []byte(*flag.String("body", "Hello geeks!!!", "send mail body"))

	flag.Parse()
	host := "smtp.mail.yahoo.com"
	port := "587"
	auth := smtp.PlainAuth("", *u, *p, host)
	fmt.Println(auth)
	for i := 0; i < *num; i++ {
		err := smtp.SendMail(host+":"+port, auth, *u, toList, body)

		// handling the errors
		if err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("end")
}
