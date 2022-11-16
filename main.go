package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	bindFlag := flag.String("bind", ":8200", "host:port to listen on")
	flag.Parse()

	ac := NewAsanaClient()

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			handleRequest(ac, w, r)
		},
	)

	srv := http.Server{
		Addr: *bindFlag,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handleRequest(ac *AsanaClient, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
		return
	}

	name := r.Form.Get("name")
	log.Print(name)

	err = ac.CreateTask(name, "<body></body>", r.Form.Get("assignee"))
	if err != nil {
		log.Print(err)
		return
	}
}
