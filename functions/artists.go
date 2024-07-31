package functions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Artist struct represents an artist with various fields
type Artist struct {
	ID              int                 `json:"id"`
	Name            string              `json:"name"`
	Image           string              `json:"image"`
	StartYear       int                 `json:"creationDate"`
	FirstAlbum      string              `json:"firstAlbum"`
	Members         []string            `json:"members"`
	LocationsURL    string              `json:"locations"`
	DatesURL        string              `json:"concertDates"`
	RelationsURL    string              `json:"relations"`
	Locations       []string            `json:"locations"`
	Dates           []string            `json:"dates"`
	DatesByLocation map[string][]string `json:"datesByLocation"`
}

// RelationsResponse struct represents the response from the relations API
type RelationsResponse struct {
	Index []RelationData `json:"index"`
}

// RelationData struct represents a relation data item
type RelationData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// FetchData is a helper function to fetch data from a URL and unmarshal it into a target interface
func FetchData(url string, target interface{}) error {
	// Send an HTTP GET request to the specified URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal the response body into the target interface
	return json.Unmarshal(body, target)
}

// GetArtistsWithDetails fetches artists data with details from the API
func GetArtistsWithDetails() ([]Artist, error) {
	// Create a slice to hold the artists data
	var artists []Artist

	// Fetch the artists data from the API
	err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		return nil, err
	}

	// Create a variable to hold the relations response
	var relationsResponse RelationsResponse

	// Fetch the relations data from the API
	err = FetchData("https://groupietrackers.herokuapp.com/api/relation", &relationsResponse)
	if err != nil {
		return nil, err
	}
	relations := relationsResponse.Index

	// Loop through each artist and populate the DatesByLocation field
	for i, artist := range artists {
		artist.DatesByLocation = make(map[string][]string)
		for _, relation := range relations {
			if artist.ID == relation.ID {
				for location, dates := range relation.DatesLocations {
					artist.Locations = append(artist.Locations, location)
					artist.Dates = append(artist.Dates, dates...)
					artist.DatesByLocation[location] = dates
				}
				break
			}
		}
		artists[i] = artist
	}

	// Return the artists data with details
	return artists, nil
}

