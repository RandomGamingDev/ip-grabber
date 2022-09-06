package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

var redirectURL string

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
	http.Redirect(w, r, redirectURL, 301)
}

func main() {
	http.HandleFunc("/", redirect)

	scanner := bufio.NewScanner(os.Stdin)
	var scannerErr error

	fmt.Printf("What port do you want to host this on? ")
	scanner.Scan()
	scannerErr = scanner.Err()
	if scannerErr != nil {
		log.Fatal(scannerErr)
	}
	port := scanner.Text()

	fmt.Printf("What website would you like to redirect to? ")
	scanner.Scan()
	scannerErr = scanner.Err()
	if scannerErr != nil {
		log.Fatal(scannerErr)
	}
	redirectURL = scanner.Text()

	fmt.Println(fmt.Sprintf("Starting the server on :%s...", port))
	httpErr := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if httpErr != nil {
		log.Fatal("ListenAndServe: ", httpErr)
	}
}
