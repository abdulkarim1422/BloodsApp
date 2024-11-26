package services

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"

	"github.com/abdulkarim1422/BloodsApp/initializers"
	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
)

func SendMatchResult(patientID int, donors []int) {
	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		fmt.Printf("Error getting patient: %v\n", err)
		return
	}

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
	for _, donor := range donorsList {
		// Create the message
		message := fmt.Sprintf("مرحبًا %s، هذه رسالة تلقائية من تطبيق BloodsApp. يوجد مريض يحتاج إلى تبرع بالدم. يرجى التواصل مع المستشفى للتبرع.", donor.LatinName)

		// Send the message
		SendWhatsappMessage(donor.PhoneNumber, message, "match")

		// Store the request
		var request models.Request
		request.PatientID = int(patient.ID)
		request.DonorID = int(donor.ID)
		request.Address.HospitalName = patient.Address.HospitalName
		request.RedCrescentCode = patient.RedCrescentCode

		repositories.CreateRequest(&request)
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
	cmd := exec.Command(initializers.PythonCaller, "scripts/whatsapp_message.py", phone, message, txt_file_name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error sending WhatsApp message: %v\n", err)
		fmt.Printf("Command output: %s\n", output)
		return
	}
	fmt.Printf("WhatsApp message sent successfully: %s\n", output)

}
