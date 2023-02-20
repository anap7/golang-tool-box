package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

/*Retorna um ponteiro de cli.App.
Cli é o pacote baixado pelo github e o App é um tipo do pacote*/
//Retornar a aplicação de linha de comando para ser executada
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search IPs and Server's Names on Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search Address IPs and Server's Names on Internet",
			Flags: flags,
			Action:searchIps,
		},
		{
			Name: "servers",
			Usage: "Search Server's Names on Internet",
			Flags: flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	/*Net é um pacote que busca ips. Retorna
	um slice de ips ou um erro*/
	ips, error := net.LookupIP(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	/*Net é um pacote que busca ips. Retorna
	um slice de ips ou um erro. LookupNS = Name Server*/
	servers, error := net.LookupNS(host)

	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}