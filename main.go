package main

import (
	"groupie-tracker/functions"
	"html/template"
	"log"
	"net/http"
)

// Declare a global variable to hold the parsed template
var tmpl *template.Template

// Main function is the entry point of the application
func main() {
	// Parse the HTML templates
	tmpl, _ = functions.ParseTemplates()

	// Serve static files (CSS, JavaScript, images, etc.) from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle incoming requests to the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Call the HandleRequest function from the functions package
		functions.HandleRequest(w, r, tmpl)
	})

	// Log a message indicating that the server is running
	log.Println("Server is running on http://localhost:8080")

	// Start the HTTP server and listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
