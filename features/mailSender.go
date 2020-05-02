package features

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"log"
)

func MailCompose(links string, emails string) {
	//sk := []string{"sakibcoolz@gmail.com","mr.sakibmulla@gmail.com"}
	from := mail.Address{"Sakib Mulla", "sakib.m@yes-info.in"}
	to   := mail.Address{"SK", emails}
	//to := []mail.Address{}
	subj := "Add admin links - ClinicalDocs"
	body := "Please submit your details at below link.\n Click at links = http://203.192.228.45:4341/#/addadmin/"+links

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k,v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	fmt.Println(message)
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "sg2plcpnl0190.prod.sin2.secureserver.net:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("","sakib.m@yes-info.in", "chingi12!@", host)

	// TLS config
	tlsconfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: host,
	}
	fmt.Println("tls config")

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
}
