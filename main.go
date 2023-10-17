package main

import (
	"fmt"

	"github.com/kubestaff/golearn/db"
	"github.com/kubestaff/golearn/html/home"
	"github.com/kubestaff/golearn/setting"
	"github.com/kubestaff/golearn/user"
	"github.com/kubestaff/golearn/video"
	"github.com/kubestaff/web-helper/server"
)

func main() {
	dbConn, err := db.ConnectToDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Migrate(dbConn)

	opts := server.Options{Port: 4567}

	s := server.NewServer(opts)

	defer s.Stop()

	s.Handle("/", home.Handle)
	s.Handle("/user", user.Handle)
	s.Handle("/changeuser", user.HandleChange)
	s.HandleJSON("/videos", video.HandleList)

	settingsHandler := setting.Handler{
		DbConn: dbConn,
	}
	s.HandleJSON("/persist-settings", settingsHandler.Persist)
	s.HandleJSON("/settings", settingsHandler.Read)
	s.HandleJSON("/delete-settings", settingsHandler.Delete)

	s.Start()
}
