package services

import (
	"time"
)

type Donor struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"first_name" form:"FirstName"`
	LastName      string    `json:"last_name,omitempty" form:"LastName"`
	PhoneNumber   string    `json:"phone_number" form:"PhoneNumber"`
	BloodType     string    `json:"bloodType" form:"BloodType"`
	BirthDate     time.Time `json:"birth_date,omitempty" form:"BirthDate" time_format:"2006-01-02"`
	Gender        string    `json:"gender" form:"Gender"`
	CarAvailable  bool      `json:"car_available,omitempty" form:"CarAvailable"`
	Paid          bool      `json:"paid,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Address       string    `json:"address" form:"Address"`
	RedTimer      time.Time `json:"red_timer"`
	PlateletTimer time.Time `json:"platelet_timer"`
	Score         int       `json:"score"`
}

type Patient struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"firstName" form:"FirstName"`
	LastName         string    `json:"lastName,omitempty" form:"LastName"`
	BloodType        string    `json:"bloodType" form:"BloodType"`
	BirthDate        time.Time `json:"birth_date,omitempty" form:"BirthDate" time_format:"2006-01-02"`
	Gender           string    `json:"gender" form:"Gender"`
	Address          string    `json:"address" form:"Address"`
	Phone            string    `json:"phone" form:"Phone"`
	Urgency          int       `json:"urgency" form:"Urgency"`
	RedQuantity      int       `json:"red-quantity" form:"RedQuantity"`
	PlateletQuantity int       `json:"platelet-quantity" form:"PlateletQuantity"`
	CarAvailable     bool      `json:"car_available" form:"CarAvailable"`
}

var donors = []Donor{
	{ID: 1, FirstName: "John Doe", LastName: "ASD", BloodType: "A+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 2, FirstName: "Jane Doe", LastName: "ASD", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 3, FirstName: "John Smith", LastName: "ASD", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 4, FirstName: "Jane Smith", LastName: "ASD", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 5, FirstName: "John Doe", LastName: "ASD", BloodType: "A+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 6, FirstName: "Jane Doe", LastName: "ASD", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 7, FirstName: "John Smith", LastName: "ASD", BloodType: "AB+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 8, FirstName: "Jane Smith", LastName: "ASD", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 9, FirstName: "John Doe", LastName: "ASD", BloodType: "A+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 10, FirstName: "Jane Doe", LastName: "ASD", BloodType: "B+", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 11, FirstName: "John Smith", LastName: "ASD", BloodType: "AB-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 12, FirstName: "Jane Smith", LastName: "ASD", BloodType: "B-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 13, FirstName: "John Doe", LastName: "ASD", BloodType: "A-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 14, FirstName: "Jane Doe", LastName: "ASD", BloodType: "AB-", Address: "123 Main St", PhoneNumber: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
}

var patients = []Patient{
	{ID: 1, FirstName: "Jane Smith", LastName: "ASD", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 1, PlateletQuantity: 1, CarAvailable: true},
	{ID: 2, FirstName: "John Smith", LastName: "ASD", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 2, PlateletQuantity: 0, CarAvailable: false},
	{ID: 3, FirstName: "Jane Doe", LastName: "ASD", BloodType: "AB+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 1, PlateletQuantity: 1, CarAvailable: true},
}
