package main
import (
	"fmt"
	"net/smtp"
	_ "time"
)

func main() {
	from := "yusyang20121223@yahoo.com"
	toList := []string{"yusy20121223@gmail.com"}
  
	host := "smtp.mail.yahoo.com"
	port := "587"
	msg := "Hello geeks!!!"
	body := []byte(msg)
	auth := smtp.PlainAuth("", "yusyang20121223@yahoo.com", "cdtuhgrbalryenwl", host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
  
    // handling the errors
    if err != nil {
        fmt.Println(err)
	return
    }
  
    fmt.Println("end")
}
