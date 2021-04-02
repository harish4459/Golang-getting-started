package main

import (
	"net/http"

	"github.com/harish4459/webservice/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
