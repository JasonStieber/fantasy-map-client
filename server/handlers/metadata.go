package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// Location represents a point of interest on the map.
type Location struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	X           float64  `json:"x"`
	Y           float64  `json:"y"`
	Tags        []string `json:"tags"`
}

// MapData holds all locations from the map.json file.
type MapData struct {
	Locations []Location `json:"locations"`
}

// LoadMapData loads and parses the map.json file.
func LoadMapData(filePath string) (*MapData, error) {
	absPath, _ := filepath.Abs(filePath)
	file, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("error reading map data file: %v", err)
	}

	var data MapData
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("error unmarshaling map data: %v", err)
	}

	return &data, nil
}

// HandleLocations serves the location data as JSON over HTTP.
func HandleLocations(w http.ResponseWriter, r *http.Request) {
	mapData, err := LoadMapData("../data/map.json") // Adjust the path as necessary
	if err != nil {
		http.Error(w, "Failed to load map data", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(mapData); err != nil {
		http.Error(w, "Failed to encode map data", http.StatusInternalServerError)
		fmt.Println(err)
	}
}
