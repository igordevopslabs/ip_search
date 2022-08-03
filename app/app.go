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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "golangexample.com",
				},
			},
			Action: buscarIps,
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
