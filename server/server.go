package server

import (
	"net/http"

	"github.com/lucas-10101/logapi/server/routes"
)

func InitServer() {
	http.ListenAndServe("127.0.0.1:2525", routes.MakeRoutes())
}
