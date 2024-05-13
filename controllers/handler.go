package controllers

import (
	"ascii-art-web/ascii-art/funcs"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type asciiart struct {
	input  string
	font   string
	Output string
}

func Handler(w http.ResponseWriter, r *http.Request) {

	homepage, err := template.ParseFiles("home.html")

	if err != nil {
		// Display custom 500 error page
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "pages/500.html")
		return
	}

	if r.URL.Path != "/ascii-art" && r.URL.Path != "/" {

		// Display custom 404 error page
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "pages/404.html")
		return
	}

	switch r.Method {
	case "GET":
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = homepage.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	case "POST":

		ascii := asciiart{}

		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error parsing form")
			return
		}

		ascii.input = r.FormValue("in")
		ascii.font = r.FormValue("banner")
		inputArray := strings.Split(ascii.input, "\n")

		ascii.Output, err = funcs.Print(inputArray, ascii.font)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "pages/500.html")
		} else if len(ascii.input) > 500 || ascii.Output == "" || ascii.Output == " " {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "pages/400.html")
		} else {
			ascii.Output = "\n" + ascii.Output
			err = homepage.Execute(w, ascii)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

	// func exportAsciiArt(w http.ResponseWriter, r *http.Request){
	// 	// Generate the ASCII art content
	// 	asciiArt := generateAsciiArt()
	
	// 	// Set the Content-Type header for the response
	// 	w.Header().Set("Content-Type", "text/plain")
	
	// 	// Set the Content-Disposition header to specify the filename and permissions
	// 	filename := "ascii_art.txt"
	// 	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"; filename*=UTF-8''%s", filename, filename))
	
	// 	// Set the Content-Length header based on the ASCII art length
	// 	contentLength := len(asciiArt)
	// 	w.Header().Set("Content-Length", fmt.Sprintf("%d", contentLength))
	
	// 	// Write the ASCII art content to the response
	// 	_, err := w.Write([]byte(asciiArt))
	// 	if err != nil {
	// 		// Handle error if unable to write response
	// 		http.Error(w, "Failed to export ASCII art", http.StatusInternalServerError)
	// 		return
	// 	}
	// }

	func DownloadHandler(w http.ResponseWriter, r *http.Request) {
		downloadButton := r.FormValue("downloadText")
		if downloadButton == "" {
			w.WriteHeader(http.StatusBadRequest)
			parsedTemplate, _ := template.ParseFiles("../../static/400.html")
			parsedTemplate.Execute(w, nil)
			return
		}
	
		w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(downloadButton))
	}
