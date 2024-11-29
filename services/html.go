package services

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/abdulkarim1422/BloodsApp/models"
	"github.com/abdulkarim1422/BloodsApp/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Main and Dashboard pages
func Main_Page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main Page"})
}

func Dashboard_Page(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{"title": "Dashboard"})
}

// Patient and Donor forms
func ShowPatientForm(c *gin.Context) {
	c.HTML(http.StatusOK, "patient_form.html", gin.H{
		"title": "Patient Registration",
	})
}

func ShowDonorForm(c *gin.Context) {
	c.HTML(http.StatusOK, "donor_form.html", gin.H{
		"title": "Donor Registration",
	})
}

// Admin forms
func ShowMatchForm(c *gin.Context) {
	patients, err := repositories.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "match.html", gin.H{
		"title":    "Match Donor and Patient",
		"patients": patients,
	})
}

func ShowOnePatientMatchForm(c *gin.Context) {
	id := c.Param("id")
	patientID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}
	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var patients []models.Patient
	patients = append(patients, *patient)
	c.HTML(http.StatusOK, "match.html", gin.H{
		"title":    "Match Donor and Patient",
		"patients": patients,
	})
}

func ShowDonationForm(c *gin.Context) {
	patients, err := repositories.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	donors, err := repositories.GetAllDonors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "donation_form.html", gin.H{
		"donors":   donors,
		"patients": patients,
	})
}

func ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

// Schedule Request
func ShowSpecialPatientForm(c *gin.Context) {
	patients, err := repositories.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "scheduale_request_form.html", gin.H{
		"title":    "تسجيل مواعيد لمضرضى خاصّين",
		"patients": patients,
	})
}

func ShowSchedualedRequestsPage(c *gin.Context) {
	requests, err := repositories.GetAllSchedualedRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "schedualed_requests.html", gin.H{
		"title":    "Special Requests",
		"requests": requests,
	})
}

// Updating the patient and donor
func ShowUpdatePatientForm(c *gin.Context) {
	id := c.Param("id")
	patientID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}
	patient, err := repositories.GetPatientByID(patientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "update_patient.html", gin.H{
		"title":   "Update Patient",
		"patient": patient,
	})
}

func ShowUpdateDonorForm(c *gin.Context) {
	id := c.Param("id")
	donorID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid donor ID"})
		return
	}
	donor, err := repositories.GetDonorByID(donorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "update_donor.html", gin.H{
		"title": "Update Donor",
		"donor": donor,
	})
}

// Patients and Donors list
func ShowPatientsPage(c *gin.Context) {
	var patients []models.Patient
	patients, err := repositories.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "patients.html", gin.H{
		"title":    "Patients List",
		"patients": patients,
	})
}

func ShowDonorsPage(c *gin.Context) {
	var donors []models.Donor
	donors, err := repositories.GetAllDonors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "donors.html", gin.H{
		"title":  "Donors List",
		"donors": donors,
	})
}

// Waiting patients
func ShowPatientsWaitingPage(c *gin.Context) {
	patients, err := repositories.CheckPatientsWaiting()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	schedualedRequests, err := repositories.GetWaitingScheduals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "waiting_patients.html", gin.H{
		"title":              "قائمة المرضى المنتظرين",
		"patients":           patients,
		"schedualedRequests": schedualedRequests,
	})
}

// Specific request --------------------------------
func ShowSpecificRequestForDonor(c *gin.Context) {
	// Get the request ID from the query string
	reqeustID := c.Param("id")
	if reqeustID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reqeustID"})

		return
	}

	// Get the token from the query string
	tokenString := c.Param("token")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Convert the request ID to an integer
	requestIDint, err := strconv.Atoi(reqeustID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}

	// Get the request by ID
	request, err := repositories.GetRequestByID(requestIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the request status to "viewed"
	err = repositories.SetMessageOpenedByDonor(requestIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the patient by ID
	patient, err := repositories.GetPatientByID(request.PatientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the donor by ID
	donor, err := repositories.GetDonorByID(request.DonorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Render the request_donor.html template
	c.HTML(http.StatusOK, "request_donor.html", gin.H{
		"title":   "طلب تبرّع بالدّم",
		"patient": patient,
		"donor":   donor,
		"request": request,
	})
}

func ShowSpecificRequestForPatient(c *gin.Context) {
	// Get the request ID from the query string
	reqeustID := c.Param("id")
	if reqeustID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reqeustID"})

		return
	}

	// Convert the request ID to an integer
	requestIDint, err := strconv.Atoi(reqeustID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}

	// Get the request by ID
	request, err := repositories.GetRequestByID(requestIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the request status to "viewed"
	err = repositories.SetMessageOpenedByPatient(requestIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the patient by ID
	patient, err := repositories.GetPatientByID(request.PatientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the donor by ID
	donor, err := repositories.GetDonorByID(request.DonorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Render the request_patient.html template
	c.HTML(http.StatusOK, "request_patient.html", gin.H{
		"title":   "طلب تبرّع بالدّم",
		"patient": patient,
		"donor":   donor,
		"request": request,
	})
}

func RenderErrorPage(c *gin.Context, status int, message string) {
	c.HTML(status, "error.html", gin.H{
		"title":   "Error",
		"message": message,
	})
}
