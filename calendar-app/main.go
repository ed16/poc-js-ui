package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type Appointment struct {
	ID       string    `json:"id"`
	Employee string    `json:"employee"`
	Doctor   string    `json:"doctor"`
	Date     time.Time `json:"date"`
}

var (
	appointments []Appointment
	mu           sync.Mutex
)

func getAppointments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	json.NewEncoder(w).Encode(appointments)
	mu.Unlock()
}

func createAppointment(w http.ResponseWriter, r *http.Request) {
	var appointment Appointment
	_ = json.NewDecoder(r.Body).Decode(&appointment)
	appointment.ID = time.Now().Format("20060102150405")
	appointment.Date, _ = time.Parse(time.RFC3339, appointment.Date.Format(time.RFC3339)) // Ensure the date is in the correct format

	mu.Lock()
	appointments = append(appointments, appointment)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointment)
}

func main() {
	appointments = append(appointments, Appointment{
		ID:       "1",
		Employee: "John Doe",
		Doctor:   "Dr. Smith",
		Date:     time.Now(),
	})

	http.HandleFunc("/appointments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getAppointments(w, r)
		} else if r.Method == http.MethodPost {
			createAppointment(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Serve static files
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
