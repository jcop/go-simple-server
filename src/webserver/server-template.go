package main

import (
	"html/template"
	"fmt"
	"log"
	"net/http"
	"time"
	"io"
)

type PageVariables struct {
	Date string
	Time string
}

func main() {

    port := ":8080"

    log.Print("starting server on port" , port)
	http.HandleFunc("/", HomePage)

	http.HandleFunc("/health-check", HealthCheckHandler)
	log.Fatal(http.ListenAndServe( port , nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()              // find the time right now
	HomePageVars := PageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

    log.Print("STATIC got request - " , now.Format("15:04:05"))

	t, err := template.ParseFiles("templates/homepage.html") //parse the html file homepage.html
	if err != nil {                                // if there is an error
        fmt.Fprintf(w, "fatal error 1")
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
        fmt.Fprintf(w, "fatal error 2")
		log.Print("template executing error: ", err) //log it
	}

}


// e.g. http.HandleFunc("/health-check", HealthCheckHandler)
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

    log.Print("test request: ")
    // A very simple health check.
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")

    // In the future we could report back on the status of our DB, or our cache
    // (e.g. Redis) by performing a simple PING, and include them in the response.
    io.WriteString(w, `{"alive": true}`)
}


