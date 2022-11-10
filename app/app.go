package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli/v2"
)

// Gerar é responsável por gerar a aplicação
func Gerar() *cli.App {
	// Criamos uma nova aplicação
	app := cli.NewApp()
	// Definimos o nome da aplicação
	app.Name = "Aplicação de Linha de Comando"
	// Definimos a descrição da aplicação
	app.Usage = "Busca IPs e Nomes de Servidores na Internet"

	// Definimos os comandos da aplicação
	app.Commands = []*cli.Command {
		{
			// Nome do comando
			Name: "ip",
			// Descrição do comando
			Usage: "Busca IPs de endereços na Internet",
			// Flags são as opções que o comando aceita
			Flags: []cli.Flag {
				// Fazemos referência a flag host e definimos o nome dela, o valor padrão e a descrição
				&cli.StringFlag{
					Name: "host",
					Value: "google.com",
					Usage: "Host a ser buscado",
				},
			},
			// Action é a função que será executada quando o comando for chamado
			Action: buscarIps,
		},
	}
	
	return app
}

func buscarIps(c *cli.Context) error {
	// Passamos a flag host para a variável host
	host := c.String("host")

	// Pacote net é responsável por realizar a busca de IPs
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	// Imprimimos os IPs
	for _, ip := range ips {
		fmt.Println(ip)
	}

	// Retornamos nil para indicar que não houve erro
	return nil
}