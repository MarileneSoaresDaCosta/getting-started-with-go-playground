package main

import (
	"net/http"
	"playground/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	// this is enough to see the message at localhost:3000/users
}
