package main

import (
	"daidai-server/pkg/setting"
	"daidai-server/routers"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

type Test struct {
	gorm.Model
	Id   uint
	Name string
}

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 2,
	}

	s.ListenAndServe()
}
