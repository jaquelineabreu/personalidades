package main

import (
	"fmt"

	"github.com/jaquelineabreu/personalidades/models"
	"github.com/jaquelineabreu/personalidades/routes"
)



func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome1", Historia:"Historia 1"},
		{Nome: "Nome2", Historia:"Historia 2"},
	}

	fmt.Println("Inciando o servidor Rest com GO")
	routes.HandleRequest()
}