package services

import (
	"fmt"
	"os"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
)

// var subject string = "Feedback"
// var body string = "Thank you for your feedback. We will get back to you soon."

// func SendFeedBackToPatient(mail string, phone string) {
// 	// Send feedback mail to the patient
// 	emailReq := EmailRequest{
// 		To:      mail,
// 		Subject: subject,
// 		Body:    body,
// 	}
// 	if err := SendMail(emailReq); err != nil {
// 		fmt.Printf("Failed to send feedback email: %v, %v\n", mail, err)
// 		return
// 	}

// 	// Send feedback message to the patient
// 	SendWhatsappMessage(phone, body, "feedback_patient")
// }

// func SendFeedBackToDonor(mail string, phone string) {
// 	// Send feedback mail to the donor
// 	emailReq := EmailRequest{
// 		To:      mail,
// 		Subject: subject,
// 		Body:    body,
// 	}
// 	if err := SendMail(emailReq); err != nil {
// 		fmt.Printf("Failed to send feedback email: %v, %v\n", mail, err)
// 		return
// 	}

// 	// Send feedback message to the donor
// 	SendWhatsappMessage(phone, body, "feedback_donor")
// }

// func SendFeedback(patient models.Patient, donor models.Donor) {
// 	SendFeedBackToPatient(patient.Email, patient.PhoneNumber)
// SendFeedBackToDonor(donor.Email, donor.PhoneNumber)
// }

// send feedback to the patient

func RecievePetientFeedbackInfo(request *models.Request) {
	patient, err := repositories.GetPatientByID(request.PatientID)
	if err != nil {
		fmt.Printf("Failed to get patient by id: %v\n", request.PatientID)
		return
	}

	donor, err := repositories.GetDonorByID(request.DonorID)
	if err != nil {
		fmt.Printf("Failed to get donor by id: %v\n", request.DonorID)
		return
	}

	// Create the message body
	var domain = os.Getenv("domain")
	message := fmt.Sprintf(
		`مرحبًا %s،
هذه رسالة تلقائية من تطبيق BloodsApp.
وفقاً للبيانات التي وصلتنا، لقد قام بالتبرّع بالدّم لك %s.
نرجو منك تأكيد الاستلام وتقييمه عبر ملء الاستمارة التالية:
%s/request_patient/%d

هذه الخطوة مهمّة لكي لا تحدث أي أخطاء. شكراً لتعاونكم.
`,
		patient.LatinName, donor.LatinName, domain, request.ID)

	SendFormToPatient(patient.Email, patient.PhoneNumber, message)

	fmt.Printf("Feedback message created to send to the patient: %v\n", patient.ID)
}

func SendFormToPatient(mail string, phone string, body string) {
	// Send feedback mail to the patient
	emailReq := EmailRequest{
		To:      mail,
		Subject: "Form for the patient",
		Body:    body,
	}
	if err := SendMail(emailReq); err != nil {
		fmt.Printf("Failed to send feedback email: %v, %v\n", mail, err)
		return
	}

	// Send feedback message to the patient
	SendWhatsappMessage(phone, body, "patient_feedback")

	fmt.Printf("Feedback message sent to the patient: %v\n", phone)
}
