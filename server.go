package main

import (
	"github.com/Arthelon/n10n/models"
	"log"
	"fmt"
	"net/http"
	"github.com/Arthelon/n10n/controllers"
)

func main() {
	err := models.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Listening on port %v\n", models.Conf.Port[1:])
	http.ListenAndServe(models.Conf.Port, controllers.GetRoutes())
}