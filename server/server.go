package server

import (
	"fmt"
	"net/http"

	"github.com/lucas-10101/logapi/server/routes"
	"github.com/lucas-10101/logapi/settings"
)

func InitServer() {

	routes := routes.MakeRoutes()

	fmt.Printf("Listening on %s:%d\n",
		settings.GetApplicationProperties().GetServerProperties().GetServerHost(),
		settings.GetApplicationProperties().GetServerProperties().GetServerPort(),
	)
	err := http.ListenAndServe(
		fmt.Sprintf("%s:%d",
			settings.GetApplicationProperties().GetServerProperties().GetServerHost(),
			settings.GetApplicationProperties().GetServerProperties().GetServerPort(),
		),
		routes,
	)

	if err != nil {
		panic(
			fmt.Sprintf("Cant initiate server on host %s and port %d",
				settings.GetApplicationProperties().GetServerProperties().GetServerHost(),
				settings.GetApplicationProperties().GetServerProperties().GetServerPort(),
			),
		)
	}
}
