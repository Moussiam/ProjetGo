package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func currentHour(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, time.Now().Format("15:04"))
}

func addToFile(author, entry string) {
	f, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(author + ": " + entry + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Enregistré dans le fichier file.txt")
}

func add(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if err := req.ParseForm(); err != nil {
		fmt.Println("Bad request")
		fmt.Fprintln(w, "Bad request")
		return
	}
	author := req.PostForm.Get("author")
	entry := req.PostForm.Get("entry")

	fmt.Println(author, ": ", entry)
	if len(author) > 0 && len(entry) > 0 {
		addToFile(author, entry)
		fmt.Fprintf(w, author+": "+entry)
	} else {
		fmt.Fprintf(w, "Missing parameters")
	}

}

func save(author, entry string) {
	f, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(author + ": " + entry + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Enregistré dans le fichier file.txt")
}

func entries(w http.ResponseWriter, req *http.Request) {
	raw, err := os.ReadFile("file.txt")

	if err != nil {
		panic(err)
	}

	entries := strings.Split(string(raw), "\n")

	for _, entry := range entries {
		entry := strings.Split(entry, ":")

		fmt.Fprintf(w, entry[0]+"\n")
	}
}

func main() {
	http.HandleFunc("/", currentHour)
	http.HandleFunc("/add", add)
	http.HandleFunc("/entries", entries)
	http.ListenAndServe(":4567", nil)
}
