package main

import (
	"log"
	"net/smtp"
)

/*
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=srv1223215@gmail.com
SMTP_PASSWORD="qinc fnoz hrif jryb"
SMTP_FROM="minhhoccode111 - eForm <srv1223215@gmail.com>"
*/

func main() {
	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", "srv1223215@gmail.com", "qinc fnoz hrif jryb", smtpHost)

	// message
	from := "srv1223215@gmail.com"
	to := []string{"minhhoccode111@gmail.com"}
	subject := "Subject: Hello from Go!\n"
	body := "This is email body tai vi sao\n"
	message := []byte(subject + "\n" + body)

	// send
	addr := smtpHost + ":" + smtpPort
	err := smtp.SendMail(addr, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Email send successfully")
}
