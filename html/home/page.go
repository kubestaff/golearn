package home

import "github.com/kubestaff/web-helper/server"

func Handle(inputs server.Input) (filename string, placeholders map[string]string) {
	return "html/index.html", nil
}
