package main

import (
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	// u := models.User{
	// 	ID:        2,
	// 	FirstName: "Trii",
	// 	LastName:  "Maa",
	// }
	// fmt.Println(u)
}
