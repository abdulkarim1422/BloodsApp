package services

import (
	"fmt"
	"os"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
)

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

	// Generate JWT
	tokenString, err := GenerateRequestJWT(int(request.ID))
	if err != nil {
		fmt.Printf("Error generating JWT: %v\n", err)
		return
	}

	// Create the message body
	var domain = os.Getenv("domain")
	message := fmt.Sprintf(
		`مرحبًا %s،
هذه رسالة تلقائية من تطبيق BloodsApp.
وفقاً للبيانات التي وصلتنا، لقد قام بالتبرّع بالدّم لك %s.
نرجو منك تأكيد الاستلام وتقييمه عبر ملء الاستمارة التالية:
%s/request_patient/%d?token=%s

هذه الخطوة مهمّة لكي لا تحدث أي أخطاء. شكراً لتعاونكم.
`,
		patient.LatinName, donor.LatinName, domain, request.ID, tokenString)

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
	SendSMSMessage(phone, body, "patient_feedback")

	fmt.Printf("Feedback message sent to the patient: %v\n", phone)
}
