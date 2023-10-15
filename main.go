package main

import (
	// "github.com/kubestaff/golearn/db"
	"github.com/kubestaff/golearn/html/home"
	"github.com/kubestaff/golearn/setting"
	"github.com/kubestaff/golearn/user"
	"github.com/kubestaff/golearn/video"
	"github.com/kubestaff/web-helper/server"
	//"fmt"
	
)

func main() {

	// dbConn, err := db.ConnectDb()
	// if err!= nil {
	// 	fmt.Println(err)
	// 	return

	// }
	
	// db.Migrate(dbConn)


	opts := server.Options{Port: 4567}
	// we create the simplified web server
	s := server.NewServer(opts)

	// we close the server at the end
	defer s.Stop()

	s.Handle("/", home.Handle)
	s.Handle("/user", user.Handle)
	s.Handle("/changeuser", user.HandleChange)
	s.HandleJSON("/videos", video.HandleList)
	s.HandleJSON("/persist-settings", setting.HandlePersist)
	s.HandleJSON("/settings", setting.HandleRead)

	// we start the webserver don't put any code after it
	s.Start()
}
