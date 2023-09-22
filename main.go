package main

import (
"github.com/kubestaff/golearn/user"
"github.com/kubestaff/web-helper/server"
)

func main() {
opts := server.Options{Port: 4567}
// we create the simplified web server
s := server.NewServer(opts)

// we close the server at the end
defer s.Stop()

userProvider := user.Provider{}

user, err := userProvider.GetCurrentUser()
if err != nil {
	variables := map[string]string{
		"%errorMessage%": err.Error(),
	}
	s.PrintFile("/errors","error.html", variables)
} else {
	variables := map[string]string{
		"%name%":            user.GetFullName(),
		"%profile-picture%": user.GetProfilePic(),
		"%profileColor%":    user.GetProfileColor(),
	}
	s.PrintFile("/", "index.html", variables)
}

// we start the webserver don't put any code after it
s.Start()
}

