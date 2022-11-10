package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando a aplicação...")
	
	// Inicia a aplicação
	aplicacao := app.Gerar()
	// O método Run() é responsável por executar a aplicação e os.Args são os argumentos que serão passados para a aplicação via linha de comando
	if err := aplicacao.Run(os.Args); err != nil {
		// O método log.Fatal() é responsável por exibir uma mensagem de erro e encerrar a aplicação
		log.Fatal(err)
	}
	
}