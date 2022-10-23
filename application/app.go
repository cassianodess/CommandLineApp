package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Fetch IP address and Server names from internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Fetch IP address from internet",
			Flags: flags,
			Action: FetchIps,
		},
		{
			Name: "servers",
			Usage: "Fetch Server Names",
			Flags: flags,
			Action: FetchServeNames,
			
		},
	}
	
	return app
}

func FetchIps(context *cli.Context) {
	host := context.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func FetchServeNames(context *cli.Context) {
	host := context.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}