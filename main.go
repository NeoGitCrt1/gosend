package main

import (
	"flag"
	"fmt"
	"net/smtp"
	"strings"
	_ "time"
)
// "cdtuhgrbalryenwl"
func main() {
	from := flag.String("from", "test@d.a", "send from")
	toList := strings.Split(*(flag.String("out", "test@d.c;test@d.d", "send to")), ";")
	num := flag.Int("send num", 1, "how many to send")
	u := flag.String("user", "", "username: xxx@yahoo.com")
	p := flag.String("app password", "", "generated app password")
	host := "smtp.mail.yahoo.com"
	port := "587"
	body := []byte(*flag.String("body", "Hello geeks!!!", "send mail body"))
	auth := smtp.PlainAuth("", *u, *p, host)
	
	for i := 0 ; i < *num; i++ {
		err := smtp.SendMail(host+":"+port, auth, *from, toList, body)

		// handling the errors
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	

	fmt.Println("end")
}
