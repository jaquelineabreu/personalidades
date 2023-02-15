package routes

import (
	"log"
	"net/http"
)

func HandleRequest(){
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
