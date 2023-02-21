package main

import (
	"fmt"

	"github.com/jaquelineabreu/personalidades/database"
	"github.com/jaquelineabreu/personalidades/routes"
)



func main() {

	database.ConectaComBancoDeDados()

	fmt.Println("Inciando o servidor Rest com GO")
	routes.HandleRequest()
}