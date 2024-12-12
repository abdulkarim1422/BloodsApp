package services

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
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

	// Run SMS script Loop
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
	%s/request_donor/%d?token=%s
	`,
			donor.LatinName, patient.BloodType, patient.Address.HospitalName, domain, requestID, tokenString)

		// Send the message
		SendSMSMessage(donor.PhoneNumber, message, "match")

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
	SendSMSMessage(phone, msg, "verification")
	fmt.Printf("Verification code sent successfully: %s\n", code)
}

func SendSMSMessage(phone string, message string, txt_file_name string) {
	encodedMessage := url.QueryEscape(message)
	url := fmt.Sprintf("%s?phone=%s&message=%s&txt_file_name=%s", os.Getenv("SMS_MS"), phone, encodedMessage, txt_file_name)
	resp, err := http.Post(url, "application/json", nil)
	fmt.Printf("sending post request: %s\n", url)
	if err != nil {
		fmt.Printf("Failed to send SMS: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to send SMS: received status code %d\n", resp.StatusCode)
		return
	}

	fmt.Printf("SMS sent successfully to %s\n", phone)
}
