package main

import (
	"fmt"

	"github.com/jaquelineabreu/personalidades/database"
	"github.com/jaquelineabreu/personalidades/models"
	"github.com/jaquelineabreu/personalidades/routes"
)



func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia:"Historia 1"},
		{Id: 2, Nome: "Nome2", Historia:"Historia 2"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Inciando o servidor Rest com GO")
	routes.HandleRequest()
}