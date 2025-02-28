package web

import (
	"fmt"
	"net/http"
)

func ServerInit() {
	err := http.ListenAndServe("127.0.0.1:80", Router())
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Server started !")
}
