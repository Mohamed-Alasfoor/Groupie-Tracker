package functions

import (
	"html/template"
	"log"
)

func ParseTemplates() (*template.Template, error) {
	// Parse the HTML template files
	tmpl, err := template.ParseFiles("templates/template.html", "templates/artist.html", "templates/error.html")
	if err != nil {
		// If there's an error parsing the templates, log a fatal error and exit the program
		log.Fatal("Error parsing templates: ", err)
	}

	// Return the parsed template and nil error if successful
	return tmpl, nil
}

