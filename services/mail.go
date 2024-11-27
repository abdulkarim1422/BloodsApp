package services

import (
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type EmailRequest struct {
	To      string `json:"to" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

func SendMail(emailReq EmailRequest) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	recipients := strings.Split(emailReq.To, ",")
	// Filter out empty email addresses
	var validRecipients []string
	for _, recipient := range recipients {
		if strings.TrimSpace(recipient) != "" {
			validRecipients = append(validRecipients, recipient)
		}
	}

	// Log the recipient email addresses
	log.Printf("Sending email to: %v", validRecipients)

	if len(validRecipients) == 0 {
		log.Println("No valid email addresses to send to.")
		return nil
	}

	message := []byte("Subject: " + emailReq.Subject + "\r\n\r\n" + emailReq.Body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, validRecipients, message)
	if err != nil {
		log.Fatalf("smtp error: %s", err)
		return err
	}

	log.Println("Mail sent successfully")
	return nil
}
