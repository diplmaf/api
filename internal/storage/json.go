package storage

import (
	"encoding/json"
	"os"
	"sync"
	"travel-api/internal/models"
)

type Storage struct {
	trips []models.Trip
	mu    sync.RWMutex
	file  string
}

func NewJSONStorage(filename string) *Storage {
	s := &Storage{file: filename}
	s.loadFromFile()
	return s
}

func (s *Storage) loadFromFile() {
	file, err := os.Open(s.file)
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		return
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&s.trips)
}

func (s *Storage) saveToFile() {
	file, _ := os.Create(s.file)
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") 
	if err := encoder.Encode(s.trips); err != nil {
    
}
}
func (s *Storage) GetAllTrips() []models.Trip {
	s.mu.RLock()
	defer s.mu.RUnlock()
	trips := make([]models.Trip, len(s.trips))
	copy(trips, s.trips)
	return trips
}

func (s *Storage) GetTripByID(id string) (*models.Trip, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i := range s.trips {
		if s.trips[i].ID == id {
			return &s.trips[i], true
		}
	}
	return nil, false
}

func (s *Storage) CreateTrip(trip models.Trip) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.trips = append(s.trips, trip)
	s.saveToFile()
}

func (s *Storage) UpdateTrip(id string, trip models.Trip) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.trips {
		if s.trips[i].ID == id {
			s.trips[i] = trip
			s.saveToFile()
			return true
		}
	}
	return false
}

func (s *Storage) DeleteTrip(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.trips {
		if s.trips[i].ID == id {
			s.trips = append(s.trips[:i], s.trips[i+1:]...)
			s.saveToFile()
			return true
		}
	}
	return false
}
