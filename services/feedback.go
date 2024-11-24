package services

import (
	"fmt"

	"github.com/abdulkarim1422/BloodsApp/models"
)

var subject string = "Feedback"
var body string = "Thank you for your feedback. We will get back to you soon."

func SendFeedBackToPatient(mail string, phone string) {
	// Send feedback mail to the patient
	emailReq := EmailRequest{
		To:      mail,
		Subject: subject,
		Body:    body,
	}
	if err := SendMail(emailReq); err != nil {
		fmt.Printf("Failed to send feedback email: %v, %v\n", mail, err)
		return
	}

	// Send feedback message to the patient
	SendWhatsappMessage(phone, body, "feedback_patient")
}

func SendFeedBackToDonor(mail string, phone string) {
	// Send feedback mail to the donor
	emailReq := EmailRequest{
		To:      mail,
		Subject: subject,
		Body:    body,
	}
	if err := SendMail(emailReq); err != nil {
		fmt.Printf("Failed to send feedback email: %v, %v\n", mail, err)
		return
	}

	// Send feedback message to the donor
	SendWhatsappMessage(phone, body, "feedback_donor")
}

// May use later
func SendFeedback(patient models.Patient, donor models.Donor) {
	SendFeedBackToPatient(patient.Email, patient.PhoneNumber)
	SendFeedBackToDonor(donor.Email, donor.PhoneNumber)
}
