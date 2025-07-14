package main

import (
	"net/http"
	"tiktok001/router"
	"tiktok001/router/routers"
	"tiktok001/setting"
)

func main() {
	setting.Init()

	r := routers.NewRouter()

	server := &http.Server{}

}
