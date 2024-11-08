package services

import "time"

type Donor struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	BloodType     string    `json:"bloodType"`
	Address       string    `json:"address"`
	Phone         string    `json:"phone"`
	RedTimer      time.Time `json:"redTimer"`
	PlateletTimer time.Time `json:"plateletTimer"`
	Score         int       `json:"score"`
}

type Patient struct {
	ID               int    `json:"id"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	BloodType        string `json:"bloodType"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	Urgency          int    `json:"urgency"`
	RedQuantity      int    `json:"red-quantity"`
	PlateletQuantity int    `json:"platelet-quantity"`
	Capability       string `json:"capability"`
}

var donors = []Donor{
	{ID: 1, FirstName: "John Doe", LastName: "ASD", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 2, FirstName: "Jane Doe", LastName: "ASD", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 3, FirstName: "John Smith", LastName: "ASD", BloodType: "AB+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 4, FirstName: "Jane Smith", LastName: "ASD", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 5, FirstName: "John Doe", LastName: "ASD", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 12, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 6, FirstName: "Jane Doe", LastName: "ASD", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 7, FirstName: "John Smith", LastName: "ASD", BloodType: "AB+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 8, FirstName: "Jane Smith", LastName: "ASD", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 9, FirstName: "John Doe", LastName: "ASD", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 10, FirstName: "Jane Doe", LastName: "ASD", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
}

var patients = []Patient{
	{ID: 1, FirstName: "Jane Smith", LastName: "ASD", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 1, PlateletQuantity: 1, Capability: "red"},
	{ID: 2, FirstName: "John Smith", LastName: "ASD", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 2, PlateletQuantity: 0, Capability: "platelet"},
	{ID: 3, FirstName: "Jane Doe", LastName: "ASD", BloodType: "AB+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 1, PlateletQuantity: 1, Capability: "red"},
}
