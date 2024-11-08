package main

import "time"

type donor struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	BloodType     string    `json:"bloodType"`
	Address       string    `json:"address"`
	Phone         string    `json:"phone"`
	RedTimer      time.Time `json:"redTimer"`
	PlateletTimer time.Time `json:"plateletTimer"`
	Score         int       `json:"score"`
}

type patient struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	BloodType        string `json:"bloodType"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	Urgency          int    `json:"urgency"`
	RedQuantity      int    `json:"red-quantity"`
	PlateletQuantity int    `json:"platelet-quantity"`
	Capability       string `json:"capability"`
}

var donors = []donor{
	{ID: 1, Name: "John Doe", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 2, Name: "Jane Doe", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 3, Name: "John Smith", BloodType: "AB+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 4, Name: "Jane Smith", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 5, Name: "John Doe", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 12, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 6, Name: "Jane Doe", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 7, Name: "John Smith", BloodType: "AB+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 8, Name: "Jane Smith", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 9, Name: "John Doe", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
	{ID: 10, Name: "Jane Doe", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", RedTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), PlateletTimer: time.Date(2024, 11, 7, 0, 0, 0, 0, time.UTC), Score: 0},
}

var patients = []patient{
	{ID: 1, Name: "Jane Smith", BloodType: "A+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 1, PlateletQuantity: 1, Capability: "red"},
	{ID: 2, Name: "John Smith", BloodType: "B+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 2, PlateletQuantity: 0, Capability: "platelet"},
	{ID: 3, Name: "Jane Doe", BloodType: "AB+", Address: "123 Main St", Phone: "555-555-5555", Urgency: 0, RedQuantity: 1, PlateletQuantity: 1, Capability: "red"},
}
