package main

import (
	"github.com/kubestaff/golearn"
)
func main() {
	opts := server.Options{}
	// we create the simplified web server
	s := server.NewServer(opts)

	// we close the server at the end
	defer s.Stop()

	variables := map[string]string{"%name%": "Tareq H"}

	// we output the contents of index.html
	s.PrintFile("index.html", variables)

	// we start the webserver don't put any code after it
	s.Start()
}