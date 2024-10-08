package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atotto/clipboard"
)

func getClipboard(w http.ResponseWriter, r *http.Request) {
	// Get clipboard content
	text, err := clipboard.ReadAll()
	if err != nil {
		http.Error(w, "Failed to read clipboard", http.StatusInternalServerError)
		return
	}

	// Send clipboard content as response
	fmt.Fprintf(w, "Clipboard content: %s", text)
}

func main() {
	http.HandleFunc("/clipboard", getClipboard)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
