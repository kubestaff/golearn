package main

import (
	"github.com/kubestaff/golearn/html/home"
	"github.com/kubestaff/golearn/setting"
	"github.com/kubestaff/golearn/user"
	"github.com/kubestaff/golearn/video"
	"github.com/kubestaff/web-helper/server"
)

func main() {
	opts := server.Options{Port: 4567}

	s := server.NewServer(opts)

	defer s.Stop()

	s.Handle("/", home.Handle)
	s.Handle("/user", user.Handle)
	s.Handle("/changeuser", user.HandleChange)
	s.HandleJSON("/videos", video.HandleList)
	s.HandleJSON("/persist-settings", setting.HandlePersist)
	s.HandleJSON("/settings", setting.HandleRead)

	s.Start()
}
