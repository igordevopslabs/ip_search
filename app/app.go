package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//A funcao Gerar vai retornar a app de linha de comando pronta pra ser usada
func Gerar() *cli.App {
	app := cli.NewApp()

	app.Name = "App de linha de comando"
	app.Usage = "Busca IPs e Dominios acessiveis da internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "golangexample.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servers",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscaServers,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	//net

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscaServers(c *cli.Context) {
	host := c.String("host")

	//net LookupNS

	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
