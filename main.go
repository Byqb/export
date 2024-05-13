package main

import (
	"ascii-art-web/controllers"
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", controllers.Handler) //call the function
	http.HandleFunc("/download", controllers.DownloadHandler)

	fmt.Println(controllers.GetIp())

	fmt.Println("Server listening on port http://localhost:8080/ascii-art ...") // serve and listen to port 8080

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error while listening and serving", err)
		os.Exit(0)
	}
	http.HandleFunc("/", controllers.Handler)
}
