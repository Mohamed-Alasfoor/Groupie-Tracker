package functions

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func HandleRequest(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	// Use a switch statement to handle different URL paths
	switch r.URL.Path {
	case "/":
		// Get the search query from the URL query parameters
		searchQuery := r.URL.Query().Get("search")

		// Fetch the artists data with details
		artists, err := GetArtistsWithDetails()
		if err != nil {
			// If there's an error fetching the data, render an error page
			RenderError(w, tmpl, "Error fetching artists data", http.StatusInternalServerError)
			return
		}

		// If a search query is provided
		if searchQuery != "" {
			// Create a new slice to store filtered artists
			var filteredArtists []Artist

			// Loop through all artists
			for _, artist := range artists {
				// Check if the artist's name (converted to lowercase) contains the search query (also in lowercase)
				if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(searchQuery)) {
					// If it does, add the artist to the filtered slice
					filteredArtists = append(filteredArtists, artist)
				}
			}

			// Replace the original artists slice with the filtered slice
			artists = filteredArtists

			// If no artists were found after filtering
			if len(artists) == 0 {
				// Render an error page indicating that no artist was found
				RenderError(w, tmpl, "No artist found by that name", http.StatusNotFound)
				return
			}
		}

		// Render the "template.html" template with the artists data
		err = tmpl.ExecuteTemplate(w, "template.html", artists)
		if err != nil {
			// If there's an error rendering the template, render an error page
			RenderError(w, tmpl, "Error rendering template", http.StatusInternalServerError)
		}
	case "/artist":
		// Get the artist ID from the URL query parameters
		id := r.URL.Query().Get("id")
		if id == "" {
			// If no artist ID is provided, render an error page
			RenderError(w, tmpl, "Artist ID is required", http.StatusBadRequest)
			return
		}

		// Fetch the artists data with details
		artists, err := GetArtistsWithDetails()
		if err != nil {
			// If there's an error fetching the data, render an error page
			RenderError(w, tmpl, "Error fetching artist data", http.StatusInternalServerError)
			return
		}

		// Create a new Artist variable to store the requested artist
		var artist Artist

		// Loop through all artists
		for _, a := range artists {
			// Check if the artist's ID matches the requested ID
			if fmt.Sprintf("%d", a.ID) == id {
				// If it does, store the artist and break out of the loop
				artist = a
				break
			}
		}

		// If no artist was found with the requested ID
		if artist.ID == 0 {
			// Render an error page indicating that the artist was not found
			RenderError(w, tmpl, "Artist not found", http.StatusNotFound)
			return
		}

		// Render the "artist.html" template with the artist data
		err = tmpl.ExecuteTemplate(w, "artist.html", artist)
		if err != nil {
			// If there's an error rendering the template, render an error page
			RenderError(w, tmpl, "Error rendering artist page", http.StatusInternalServerError)
		}
	default:
		// If the URL path doesn't match any of the cases above, render a "Page not found" error page
		RenderError(w, tmpl, "Page not found", http.StatusNotFound)
	}
}

