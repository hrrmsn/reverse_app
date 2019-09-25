package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	APP_NAME   = "reverse"
	ERROR_WORD = "error"
	ERROR_TEXT = "Error has occurred."
)

var workingDir string

func init() {
	var err error
	workingDir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}

func reverser(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Path[1:]
	if strings.ToLower(word) == ERROR_WORD {
		w.Write([]byte(ERROR_TEXT))
		return
	}
	output, err := exec.Command(workingDir+"/"+APP_NAME, word).Output()
	if err != nil {
		log.Fatal(err)
	}
	w.Write(output)
}

func main() {
	http.HandleFunc("/", reverser)
	fmt.Println("listening on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
