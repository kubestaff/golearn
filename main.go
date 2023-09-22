package main

import (
	"github.com/kubestaff/golearn/html/home"
	"github.com/kubestaff/golearn/user"
	"github.com/kubestaff/web-helper/server"
)

func main() {
	opts := server.Options{Port: 4567}
	// we create the simplified web server
	s := server.NewServer(opts)

	// we close the server at the end
	defer s.Stop()

	s.Handle("/", home.Handle)
	s.Handle("/user", user.Handle)

	// we start the webserver don't put any code after it
	s.Start()
}

