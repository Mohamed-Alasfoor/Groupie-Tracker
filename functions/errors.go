package functions

import (
	"net/http"
	"html/template"
)

func RenderError(w http.ResponseWriter, tmpl *template.Template, message string, code int) {
	// Set the HTTP status code for the response
	w.WriteHeader(code)

	// Create an anonymous struct to hold the error message and code
	err := tmpl.ExecuteTemplate(w, "error.html", struct {
		Message string
		Code    int
	}{
		Message: message, // Set the error message
		Code:    code,    // Set the error code
	})

	// If there's an error rendering the error page
	if err != nil {
		// Write an error message to the response with a 500 Internal Server Error status code
		http.Error(w, "Error rendering error page", http.StatusInternalServerError)
	}
}

