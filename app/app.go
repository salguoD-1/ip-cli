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

	// flags são as opções que o comando aceita
	flags := []cli.Flag {
		// Fazemos referência a flag host e definimos o nome dela, o valor padrão e a descrição
		&cli.StringFlag{
			Name: "host",
			Value: "google.com",
			Usage: "Host a ser buscado",
		},
	}

	// Definimos os comandos da aplicação
	app.Commands = []*cli.Command {
		{
			// Nome do comando
			Name: "ip",
			// Descrição do comando
			Usage: "Busca IPs de endereços na Internet",
			// Passamos a variável flags para o comando
			Flags: flags,
			// Action é a função que será executada quando o comando buscarIps for chamado
			Action: buscarIps,
		},

		{
			// Nome do comando
			Name: "servidores",
			// Descrição do comando
			Usage: "Busca o nome dos servidores na internet",
			// Passamos a variável flags para o comando
			Flags: flags,
			// Action é a função que será executada quando o comando buscarServidores for chamado
			Action: buscarServidores,
		},
	}
	// Retornamos a aplicação
	return app
}

// Função buscarIps é responsável por buscar os IPs de um determinado host
func buscarIps(c *cli.Context) error {
	// Passamos a flag host para a variável host
	host := c.String("host")

	// A função LookupIP() retorna um slice de IPs
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

// Função buscarServidores é responsável por buscar o nome dos servidores de um determinado host
func buscarServidores(c *cli.Context) error {
	// Passamos a flag host para a variável host
	host := c.String("host")

	// NS = Name Server
	// Busca o nome dos servidores de um determinado host
	servidores, err := net.LookupNS(host)
	// Tratamos o erro
	if err != nil {
		log.Fatal(err)
	}

	// Imprimimos os servidores
	for _, servidor := range servidores {
		// Usamos o .Host para não imprimir em formato struct. Ex: {ns1.google.com.}
		fmt.Println(servidor.Host)
	}

	// Retornamos nil para indicar que não houve erro
	return nil
}