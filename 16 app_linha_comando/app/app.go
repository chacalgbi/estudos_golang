package app

import (
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Aplicação para exemplificar a criação de linha de comando"
	app.Version = "1.0.0"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de endereços de IP",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) error {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		return erro
	}
	for _, ip := range ips {
		println(ip.String())
	}
	return nil
}

func buscarServidores(c *cli.Context) error {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		return erro
	}
	for _, servidor := range servidores {
		println(servidor.Host)
	}
	return nil
}
