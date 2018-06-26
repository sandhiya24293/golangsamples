/*
package main

import (
	"log"
	"net/smtp"
)

func SendMail(identity, username, password, host, message string){
	// Set up authentication information.
	auth := smtp.PlainAuth(identity, username, password, host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"recipient@example.net"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func main(){
	SendMail("", "user@example.com", "password", "mail.example.com", "This is dummy message")
}

*/

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/smtp"
	"strconv"

)
var(
	smtpHost string = "smtp.gmail.com"
	port int = 587
	username string = "donotreplyclassec@gmail.com"
	password string = "classec@123"
	to string =  "navindrakumar29@gmail.com"
)


func main() {
	// Set up authentication information.
	//auth := smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")
	auth := smtp.PlainAuth("", username, password, smtpHost)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"donotreplyclassec@gmail.com","navindrakumar29@gmail.com","donotreplyclassec@gmail.com"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: Discount Gophers!\r\n" +
		"\r\n" +
		"This is the email bodyNavindra.\r\n")
	err := smtp.SendMail(smtpHost+":"+ strconv.Itoa(port), auth, username, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
