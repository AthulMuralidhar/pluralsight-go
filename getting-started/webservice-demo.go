package main

import (
	"net/http"
	"pluralsight-go/getting-started/controllers"
)

func WebService() {

	// u1 := models.User{
	// 	Id: 2,
	// 	Firstname: "adasdad",
	// 	Lastname: "qweqwe",
	// }

	// fmt.Println(u1)
	// runtime.Breakpoint()

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
