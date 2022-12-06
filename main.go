package main

import (
	"github.com/google/martian/v3/log"
	"snake-and-ladder/application/http"
	"snake-and-ladder/conf"
	"snake-and-ladder/dao"
)

func main() {

	err := conf.Init()
	if err != nil {
		log.Errorf("conf init failed: %v", err)
		return
	}

	err = dao.Init()
	if err != nil {
		log.Errorf("dao init failed: %v", err)
		return
	}

	http.RunHTTPServer()
}
