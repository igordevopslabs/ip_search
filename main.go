package main

import (
	"ip_search/app"
	"log"
	"os"
)

func main() {

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}
