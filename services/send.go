package services

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
)

func SendMatchResult(patientID int, donors []int) {
	// Get the patient
	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		fmt.Printf("Error getting patient: %v\n", err)
		return
	}

	// Get the donors
	var donorsList []models.Donor
	for _, donorID := range donors {
		donor, err := repositories.GetDonorByID(donorID)
		if err != nil {
			fmt.Printf("Error getting donor: %v\n", err)
			return
		}
		donorsList = append(donorsList, *donor)
	}

	// Run WhatsApp script Loop
	var domain = os.Getenv("domain")
	for _, donor := range donorsList {
		// Store the request
		var request models.Request
		request.PatientID = int(patient.ID)
		request.DonorID = int(donor.ID)

		requestID, err := repositories.CreateRequest(&request)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			return
		}

		// Generate JWT
		tokenString, err := GenerateRequestJWT(requestID)
		if err != nil {
			fmt.Printf("Error generating JWT: %v\n", err)
			return
		}

		// Create the message
		message := fmt.Sprintf(
			`مرحبًا %s،
	هذه رسالة تلقائية من تطبيق BloodsApp.
	يوجد مريض يحتاج إلى تبرّع بالدم.
	زمرة الدم: %s
	المشفى: %s

	يمكنك الضغط على هذا الرابط للوصول لمعلومات المريض إذا أردت:
	%s/request_donor/%d/%s
	`,
			donor.LatinName, patient.BloodType, patient.Address.HospitalName, domain, requestID, tokenString)

		// Send the message
		SendWhatsappMessage(donor.PhoneNumber, message, "match")

		// Add one request to the patient's requests
		err_adding_request := repositories.AddOneRequestSent(patientID)
		if err_adding_request != nil {
			fmt.Printf("Error adding request to patient's requests: %v\n", err_adding_request)
			return
		}

	}

	fmt.Printf("Match results sent successfully, and stored in the Request table in database\n")
}

func generateVerificationCode() string {
	rand.Seed(int64(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func sendVerificationCode(phone, code string) {
	msg := fmt.Sprintf("رسالة التحقّق هي: %v", code)
	SendWhatsappMessage(phone, msg, "verification")
	fmt.Printf("Verification code sent successfully: %s\n", code)
}

func SendWhatsappMessage(phone string, message string, txt_file_name string) {
	pythonCaller := os.Getenv("PythonCaller")
	if pythonCaller == "" {
		pythonCaller = "python3"
	}
	cmd := exec.Command(pythonCaller, "scripts/whatsapp_message.py", phone, message, txt_file_name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error sending WhatsApp message: %v\n", err)
		fmt.Printf("Command output: %s\n", output)
		return
	}
	fmt.Printf("WhatsApp message sent successfully: %s\n", output)

}
